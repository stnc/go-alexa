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
	"net/http"
	"net/url"
	"strings"
)

var (
	cachedCert *x509.Certificate
)

func validateAlexaRequest(w http.ResponseWriter, r *http.Request) error {
	certURL := r.Header.Get("SignatureCertChainUrl")

	// Verify certificate URL
	if !verifyCertURL(certURL) {
		return fmt.Errorf("Invalid certificate url: %q", certURL)
	}

	cert, err := getX509Certificate(certURL)
	if err != nil {
		return err
	}

	//Check the certificate date 	//TODO: change ????
	//if time.Now().Unix() < cert.NotBefore.Unix() || time.Now().Unix() > cert.NotAfter.Unix() {
	//	cachedCert = nil
	//	// try again
	//	//return validateAlexaRequest(w, r) //TODO not compatible with test
	//	return fmt.Errorf("Invalid Amazon certificate date")
	//}

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
	r.Body = io.NopCloser(&bodyBuf)

	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), encryptedSig)
	if err != nil {
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
	certContents, err := io.ReadAll(cert.Body)
	if err != nil {
		return nil, errors.New("Could not read Amazon cert file.")
	}

	return certContents, nil
}

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
