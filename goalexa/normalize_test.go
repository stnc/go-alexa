package goalexa

import (
	"strings"
	"testing"
)

func Test_EscapeSSMLText(t *testing.T) {

	text1 := "Lorem &ipsum dolor & sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"
	text2 := "Lore\"m ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"
	text3 := "Lorem 'ipsum' dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"
	text4 := `<speak> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
             labor et dolore magna aliqua. In ante metus dictum at. Scelerisque purus semper
             eget duis at tellus at urna condimentum.<speak> `

	text1Convert := "&amp;"
	text2Convert := "&quot;"
	text3Convert := "&apos;"
	text4Convert := "&lt;"
	text5Convert := "&gt;"

	primeTests := []struct {
		status   string
		text     string
		getText  string
		convert  string
		expected bool
		msg      string
	}{
		{"valid", text1, "&", text1Convert, true, "expected : " + text1Convert},
		{"valid", text2, "\"", text2Convert, true, "expected : " + text2Convert},
		{"valid", text3, "'", text3Convert, true, "expected : " + text3Convert},
		{"valid", text4, "<", text4Convert, true, "expected :" + text4Convert},
		{"valid", text4, ">", text5Convert, true, "expected :" + text5Convert},

		{"invalid", text1, "&", "=", false, "expected : " + text1Convert},
		{"invalid", text2, "\"", "*", false, "expected : " + text2Convert},
		{"invalid", text3, "'", "$", false, "expected : " + text3Convert},
		{"invalid", text4, "<", "@", false, "expected :" + text4Convert},
		{"invalid", text4, ">", "-", false, "expected :" + text5Convert},
	}

	for _, e := range primeTests {
		result := EscapeSSMLText(e.text)
		contains := strings.Contains(result, e.convert)
		if e.expected && !contains {
			t.Errorf("%s: expected valid but got invalid", e.status)
		}

		if !e.expected && contains {
			t.Errorf("%s: expected invalid but got valid", e.status)
		}
	}

}
