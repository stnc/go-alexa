package goalexa

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
	validSignature := `jsHzkhi2zPaFXV4gnHN4foePDtv4SqmreEDqKqJc8kUX7skhOlZ03uKYeqLOHAot98tVJc9pMdi1TRMnkQ8sr/GoR
eO++yGi3iAYjO8/XXL1oscx1vMUzmOLmvCO/EfF3/iEpNOb3BIJEiNhT2ZIwp7EisQi3eYLDmDaklSmPWWGVQRtcSq1EoHarMW9GrUaApu2cJdAjn
F1aF3yFoLiHheN4DSW0qQ14N+ndba4C+YQBn4Ds2SXCFUyEC+q/H4A7SFioAE/qR3WYIMMfKuk1iEQOQY7jAFCS8zOjCaa4sM373T4mNUAojcgdA
aHxzu2smLRzQSttTXfuemCijTigg==`
	req.Header.Add("signaturecertchainurl", "https://s3.amazonaws.com/echo.api/echo-api-cert-7.pem")
	req.Header.Add("signature", validSignature)
	w := httptest.NewRecorder()
	errValidate := validateAlexaRequest(w, req)

	if errValidate != nil {
		t.Errorf("expected error to be nil got %v", errValidate)
	}

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
