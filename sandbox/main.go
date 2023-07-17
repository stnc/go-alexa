package main

//func main() {
//	http.HandleFunc("/upper", goalexa.ValidateAlexaRequest)
//	log.Fatal(http.ListenAndServe(":1234", nil))
//}

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

func EscapeSSMLText(text string) string {
	text = strings.ReplaceAll(text, "&", "&amp;")
	text = strings.ReplaceAll(text, "\"", "&quot;")
	text = strings.ReplaceAll(text, "'", "&apos;")
	text = strings.ReplaceAll(text, "<", "&lt;")
	text = strings.ReplaceAll(text, ">", "&gt;")
	return text
}

func main() {
	text1 := `<speak> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
             labor et dolore magna aliqua. In ante metus dictum at. Scelerisque purus semper
             eget duis at tellus at urna condimentum.<speak> `

	text2 := "Lore\"m ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"
	text3 := "Lorem 'ipsum' dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"

	fmt.Println(EscapeSSMLText(text1))
	fmt.Println(EscapeSSMLText(text2))
	fmt.Println(EscapeSSMLText(text3))
	return
	// cert, _ := getX509Certificate("https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")
	// publicKey := cert.PublicKey

	// empJSON, err := json.MarshalIndent(publicKey, "", "  ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ssl", ServeHTTP)
	http.ListenAndServe(":8090", nil)
	// fmt.Println(cert)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	certURL := r.Header.Get("SignatureCertChainUrl")

	// Verify certificate URL
	if !verifyCertURL(certURL) {

		fmt.Println("Invalid certificate url: %q", certURL)
	}
	cert, err := getX509Certificate(certURL)
	if err != nil {
		//return err
	}
	//cert, err := getX509Certificate(certURL)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	// Check the certificate date
	//if time.Now().Unix() < cert.NotBefore.Unix() || time.Now().Unix() > cert.NotAfter.Unix() {
	//	cachedCert = nil
	//	// try again
	//	return ValidateAlexaRequest(w, r)
	//}
	signa := []byte(r.Header.Get("Signature"))

	//empJSON3, err := json.MarshalIndent(signa, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	fmt.Printf("orgn output\n %s\n", string(signa))
	encryptedSigByte := base64.StdEncoding.EncodeToString(signa)

	fmt.Printf("base64  output\n %s\n", encryptedSigByte)
	// Check the certificate date
	if time.Now().Unix() < cert.NotBefore.Unix() || time.Now().Unix() > cert.NotAfter.Unix() {
		cachedCert = nil
		// try again
		fmt.Println("time-- error")
		//return validateAlexaRequest(w, r)
	}

	//openssl dgst -sha1 -sign echo-api-cert-7.pem -out filename.sha1 p.txt
	// Verify the key
	publicKey := cert.PublicKey
	encryptedSig, _ := base64.StdEncoding.DecodeString(r.Header.Get("Signature"))
	fmt.Println(encryptedSig)
	// Make the request body SHA1 and verify the request with the public key
	var bodyBuf bytes.Buffer
	hash := sha1.New()
	_, err = io.Copy(hash, io.TeeReader(r.Body, &bodyBuf))
	if err != nil {
		fmt.Println(err)
	}
	r.Body = ioutil.NopCloser(&bodyBuf)
	fmt.Printf("r.Body  output\n %s\n", r.Body)
	fmt.Printf("&bodyBuf  output\n %s\n", &bodyBuf)
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), encryptedSig)
	if err != nil {
		fmt.Println("Signature match failed")
	} else {
		fmt.Println("valid")
	}

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

	// empJSON2, err := json.MarshalIndent(certContents, "", "  ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("certContents output\n %s\n", string(empJSON2))

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
