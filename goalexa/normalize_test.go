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

	text1Expected := "&amp;"
	text2Expected := "&quot;"
	text3Expected := "&apos;"
	text4Expected := "&lt;"
	text5Expected := "&gt;"

	EscapeSSMLTextTestCases := []struct {
		status   string
		text     string
		getText  string
		convert  string
		expected bool
		msg      string
	}{
		{"valid", text1, "&", text1Expected, true, "expected : " + text1Expected},
		{"valid", text2, "\"", text2Expected, true, "expected : " + text2Expected},
		{"valid", text3, "'", text3Expected, true, "expected : " + text3Expected},
		{"valid", text4, "<", text4Expected, true, "expected :" + text4Expected},
		{"valid", text4, ">", text5Expected, true, "expected :" + text5Expected},

		{"invalid", text1, "&", "=", false, "expected : " + text1Expected},
		{"invalid", text2, "\"", "*", false, "expected : " + text2Expected},
		{"invalid", text3, "'", "$", false, "expected : " + text3Expected},
		{"invalid", text4, "<", "@", false, "expected :" + text4Expected},
		{"invalid", text4, ">", "-", false, "expected :" + text5Expected},
	}

	for _, e := range EscapeSSMLTextTestCases {
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
