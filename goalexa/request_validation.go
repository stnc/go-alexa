package goalexa

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	cachedCert *x509.Certificate
)

func StncValidateAlexaRequest(w http.ResponseWriter, r *http.Request) {
	certURL := r.Header.Get("SignatureCertChainUrl")

	// Verify certificate URL
	if !verifyCertURL(certURL) {
		//fmt.Println("Invalid certificate ")
		//return fmt.Errorf("Invalid certificate url: %q", certURL)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, certURL)
		return
	}

	cert, err := getX509Certificate(certURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "certifai err")
		return
	}

	//// Check the certificate date
	//if time.Now().Unix() < cert.NotBefore.Unix() || time.Now().Unix() > cert.NotAfter.Unix() {
	//	cachedCert = nil
	//	// try again
	//	//fmt.Println("time-- error")
	//	//return validateAlexaRequest(w, r)
	//	w.WriteHeader(http.StatusBadRequest)
	//	fmt.Fprintf(w, "time-- error")
	//	return
	//}

	// Verify the key
	publicKey := cert.PublicKey
	encryptedSig, _ := base64.StdEncoding.DecodeString(r.Header.Get("Signature"))

	// Make the request body SHA1 and verify the request with the public key
	var bodyBuf bytes.Buffer
	hash := sha1.New()
	_, err = io.Copy(hash, io.TeeReader(r.Body, &bodyBuf))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bugffer ")
		return
	}
	r.Body = ioutil.NopCloser(&bodyBuf)

	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), encryptedSig)
	if err != nil {
		//fmt.Println("Invalid Amazon certificate signature")
		//return fmt.Errorf("Invalid Amazon certificate signature: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid Amazon certificate signature")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strings.ToUpper("valid"))
}

func validateAlexaRequest(w http.ResponseWriter, r *http.Request) error {
	certURL := r.Header.Get("SignatureCertChainUrl")

	// Verify certificate URL
	if !verifyCertURL(certURL) {
		//fmt.Println("Invalid certificate ")
		return fmt.Errorf("Invalid certificate url: %q", certURL)
	}

	cert, err := getX509Certificate(certURL)
	if err != nil {
		return err
	}

	// Check the certificate date
	if time.Now().Unix() < cert.NotBefore.Unix() || time.Now().Unix() > cert.NotAfter.Unix() {
		cachedCert = nil
		// try again
		//fmt.Println("time-- error")
		return validateAlexaRequest(w, r)
	}

	// Verify the key
	publicKey := cert.PublicKey
	encryptedSig, _ := base64.StdEncoding.DecodeString(r.Header.Get("Signature"))

	// Make the request body SHA1 and verify the request with the public key
	var bodyBuf bytes.Buffer
	hash := sha1.New()
	_, err = io.Copy(hash, io.TeeReader(r.Body, &bodyBuf))
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&bodyBuf)

	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), encryptedSig)
	if err != nil {
		//fmt.Println("Invalid Amazon certificate signature")
		return fmt.Errorf("Invalid Amazon certificate signature: %v", err)
	}

	return nil
}

func getX509Certificate(certURL string) (*x509.Certificate, error) {
	if cachedCert != nil {
		return cachedCert, nil
	}

	// Fetch certificate data
	certContents, err := downloadCert(certURL)
	if err != nil {
		return nil, err
	}

	// Decode certificate data
	block, _ := pem.Decode(certContents)
	if block == nil {
		return nil, fmt.Errorf("Failed to parse Amazon certificate, %q", certURL)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Check the certificate alternate names
	foundName := false
	for _, altName := range cert.Subject.Names {
		if altName.Value == "echo-api.amazon.com" {
			foundName = true
		}
	}

	if !foundName {
		return nil, fmt.Errorf("Invalid Amazon certificate (echo-api SN not found), %q", certURL)
	}

	cachedCert = cert

	return cert, nil
}

func downloadCert(certURL string) ([]byte, error) {
	cert, err := http.Get(certURL)
	if err != nil {
		return nil, errors.New("Could not download Amazon cert file.")
	}
	defer cert.Body.Close()
	certContents, err := ioutil.ReadAll(cert.Body)
	if err != nil {
		return nil, errors.New("Could not read Amazon cert file.")
	}

	return certContents, nil
}

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-signature
func verifyCertURL(path string) bool {
	link, _ := url.Parse(path)

	if link.Scheme != "https" {
		return false
	}

	if link.Host != "s3.amazonaws.com" && link.Host != "s3.amazonaws.com:443" {
		return false
	}

	if !strings.HasPrefix(link.Path, "/echo.api/") {
		return false
	}

	return true
}
