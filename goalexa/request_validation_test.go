package mygoalexa

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-signature

func Test_validateAlexaRequest(t *testing.T) {

	jsonData, err := os.ReadFile("../mocks/requestEnvelope_SHA1.json")
	if err != nil {
		panic(err)
	}
	jsonDataStr := string(jsonData)
	payload := strings.NewReader(jsonDataStr)
	url := "http://localhost:9095/alexaTest"
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	validSignature := `jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoR
	//eO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjn
	//F1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdA
	//aHxzu2smLRzQSttTXfuemCijTigg==`
	//	req.Header.Add("signaturecertchainurl", "https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")
	//	req.Header.Add("signature", validSignature)
	//	w := httptest.NewRecorder()
	//	errValidate := validateAlexaRequest(w, req)
	//	//// iteration for solition look my testSuite 19
	//	//https://github.com/golang/go/issues/12262  https://www.google.com/search?q=golang+httptest+request+close+time&client=firefox-b-1-d&ei=QF3PZObGN-PAkPIP77u-sAQ&oq=golang+httptest+request+close&gs_lp=Egxnd3Mtd2l6LXNlcnAiHWdvbGFuZyBodHRwdGVzdCByZXF1ZXN0IGNsb3NlKgIIADIFECEYoAEyBRAhGKABMgUQIRigAUjJUFDgQVjXRnABeAGQAQCYAXGgAd8DqgEDNC4xuAEDyAEA-AEBwgIKEAAYRxjWBBiwA8ICBRAhGKsCwgIIECEYFhgeGB3iAwQYACBBiAYBkAYI&sclient=gws-wiz-serp
	//	if errValidate != nil {
	//		t.Errorf("expected error to be nil got %v", errValidate)
	//	}

	t.Run("valid signature", func(t *testing.T) {
		validSignature := `jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoR
eO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjn
F1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdA
aHxzu2smLRzQSttTXfuemCijTigg==`
		req.Header.Add("signaturecertchainurl", "https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")
		req.Header.Add("signature", validSignature)
		w := httptest.NewRecorder()
		errValidate := validateAlexaRequest(w, req)
		//// iteration for solition look my testSuite 19
		//https://github.com/golang/go/issues/12262  https://www.google.com/search?q=golang+httptest+request+close+time&client=firefox-b-1-d&ei=QF3PZObGN-PAkPIP77u-sAQ&oq=golang+httptest+request+close&gs_lp=Egxnd3Mtd2l6LXNlcnAiHWdvbGFuZyBodHRwdGVzdCByZXF1ZXN0IGNsb3NlKgIIADIFECEYoAEyBRAhGKABMgUQIRigAUjJUFDgQVjXRnABeAGQAQCYAXGgAd8DqgEDNC4xuAEDyAEA-AEBwgIKEAAYRxjWBBiwA8ICBRAhGKsCwgIIECEYFhgeGB3iAwQYACBBiAYBkAYI&sclient=gws-wiz-serp
		if errValidate != nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})

	t.Run("Invalid certificate url", func(t *testing.T) {
		validSignature := `jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoR
eO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjn
F1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdA
aHxzu2smLRzQSttTXfuemCijTigg==`
		req.Header.Add("signaturecertchainurl", "https://s3.amazonaws.com/echo.api/echo-api-cert-7.pemInvalid")
		req.Header.Add("signature", validSignature)
		w := httptest.NewRecorder()
		errValidate := validateAlexaRequest(w, req)
		if errValidate != nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})
	//TODO: change ????
	//	t.Run("Invalid Amazon certificate date,", func(t *testing.T) {
	//		validSignature := `jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoR
	//eO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjn
	//F1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdA
	//aHxzu2smLRzQSttTXfuemCijTigg==`
	//		req.Header.Add("signaturecertchainurl", "https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")
	//		req.Header.Add("signature", validSignature)
	//		w := httptest.NewRecorder()
	//		errValidate := validateAlexaRequest(w, req)
	//		if errValidate == nil {
	//			t.Errorf("expected error to be nil got %v", errValidate)
	//		}
	//	})

	t.Run("Failed to parse Amazon certificate", func(t *testing.T) {
		validSignature := `jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoR
eO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjn
F1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdA
aHxzu2smLRzQSttTXfuemCijTigg==`
		AmazonCertificateUrl := "https://s3.amazonaws.com/echo.api/echo-api-cert-7.peme"
		req.Header.Add("signaturecertchainurl", AmazonCertificateUrl)
		req.Header.Add("signature", validSignature)
		w := httptest.NewRecorder()
		errValidate := validateAlexaRequest(w, req)
		if errValidate != nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})

	t.Run("Invalid Amazon certificate (echo-api SN not found)", func(t *testing.T) {
		validSignature := `FAKEjsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoR
eO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjn
F1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdA
aHxzu2smLRzQSttTXfuemCijTigg==`
		req.Header.Add("signaturecertchainurl", "https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")
		req.Header.Add("signature", validSignature)
		w := httptest.NewRecorder()
		errValidate := validateAlexaRequest(w, req)

		if errValidate == nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})

}
func Test_getX509Certificate(t *testing.T) {

	t.Run("", func(t *testing.T) {
		_, errValidate := getX509Certificate("https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")

		if errValidate != nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})

	t.Run("Failed to parse Amazon certificate,", func(t *testing.T) {
		_, errValidate := getX509Certificate("https://s3.amazonaws.com/echo.api/echo-api-cert-7.pemFAKE")

		if errValidate != nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})

}
func Test_verifyCertURL(t *testing.T) {
	//https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-signature
	primeTests := []struct {
		name     string
		UrlName  string
		expected bool
		msg      string
	}{
		{"valid", "https://s3.amazonaws.com/echo.api/echo-api-cert.pem", true, "pass valid url"},
		{"valid", "https://s3.amazonaws.com:443/echo.api/echo-api-cert.pem", true, "pass valid url with port"},
		{"valid", "https://s3.amazonaws.com/echo.api/../echo.api/echo-api-cert.pem", true, "pass valid url with dot"},
		{"invalid", "http://s3.amazonaws.com/echo.api/echo-api-cert.pem", false, "Cannot pass url With Invalid Protocol"},
		{"invalid", "https://notamazon.com/echo.api/echo-api-cert.pem", false, "Cannot pass url With Invalid HostName"},
		{"invalid", "https://s3.amazonaws.com/EcHo.aPi/echo-api-cert.pem", false, "Cannot pass url With Invalid Path"},
		{"invalid", "https://s3.amazonaws.com:563/echo.api/echo-api-cert.pem", false, "Cannot pass url With Invalid Port"},
	}
	for _, e := range primeTests {
		result := verifyCertURL(e.UrlName)
		if e.expected && !result {
			t.Errorf("%s: expected valid but got invalid", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected invalid but got valid", e.name)
		}

	}
}
func Test_downloadCert(t *testing.T) {
	t.Run("Could not download Amazon cert file.", func(t *testing.T) {
		_, errValidate := downloadCert("https://dfdfdfdfddfdf.com/")
		fmt.Println(Convert2Err(errValidate))
		if errValidate != nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})

	t.Run("Could not read Amazon cert file.", func(t *testing.T) {
		_, errValidate := downloadCert("https://raw.githubusercontent.com/robertdavidgraham/pemcrack/master/test.dict")
		fmt.Println(Convert2Err(errValidate))
		if errValidate != nil {
			t.Errorf("expected error to be nil got %v", errValidate)
		}
	})
}

// Convert2Err error to string
func Convert2Err(err error) string {
	byteData := []byte(fmt.Sprintf("%v", err))
	return string(byteData)
}
