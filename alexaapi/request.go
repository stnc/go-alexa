package alexaapi

import (
	"context"
	"encoding/json"

	"github.com/tidwall/gjson"
)

// All possible locale codes as defined in
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/develop-skills-in-multiple-languages.html
type RequestLocale string

const (
	RequestLocaleArSA RequestLocale = "ar-SA"
	RequestLocaleDeDE RequestLocale = "de-DE"
	RequestLocaleEnAU RequestLocale = "en-AU"
	RequestLocaleEnCA RequestLocale = "en-CA"
	RequestLocaleEnGB RequestLocale = "en-GB"
	RequestLocaleEnIN RequestLocale = "en-IN"
	RequestLocaleEnUS RequestLocale = "en-US"
	RequestLocaleEsES RequestLocale = "es-ES"
	RequestLocaleEsMX RequestLocale = "es-MX"
	RequestLocaleEsUS RequestLocale = "es-US"
	RequestLocaleFrCA RequestLocale = "fr-CA"
	RequestLocaleFrFR RequestLocale = "fr-FR"
	RequestLocaleHiIN RequestLocale = "hi-IN"
	RequestLocaleItIT RequestLocale = "it-IT"
	RequestLocaleJaJP RequestLocale = "ja-JP"
	RequestLocalePtBR RequestLocale = "pt-BR"
)

type Request interface {
	GetType() RequestType
	GetRequestId() string
	GetTimestamp() string
	GetLocale() RequestLocale
	GetRequestJson() []byte
}

type RequestType string

// These are request types that don't belong to a specific interface.
//
// Other requests that DO belong to a specific interface are defined in their
// respective "iface_" files.
const (
	RequestTypeUnspecified             RequestType = ""
	RequestTypeCanFulfillIntentRequest RequestType = "CanFulfillIntentRequest"
	RequestTypeLaunchRequest           RequestType = "LaunchRequest"
	RequestTypeIntentRequest           RequestType = "IntentRequest"
	RequestTypeSessionEndedRequest     RequestType = "SessionEndedRequest"
	RequestTypeSessionResumedRequest   RequestType = "SessionResumedRequest"
)

type RequestCommon struct {
	Type      RequestType   `json:"type"`
	RequestId string        `json:"requestId"`
	Timestamp string        `json:"timestamp"`
	Locale    RequestLocale `json:"locale"`

	// Raw request JSON for custom parsing of unusual request types
	requestJson []byte
}

func (rc *RequestCommon) GetType() RequestType {
	return rc.Type
}

func (rc *RequestCommon) GetRequestId() string {
	return rc.RequestId
}

func (rc *RequestCommon) GetTimestamp() string {
	return rc.Timestamp
}

func (rc *RequestCommon) GetLocale() RequestLocale {
	return rc.Locale
}

func (rc *RequestCommon) GetRequestJson() []byte {
	return rc.requestJson
}

// Attempts to set a strongly-typed value into the root.Request field
// by "looking ahead" at the contents of the "type" field.
//
// If no match is found, a simple RequestCommon is used, and the raw request
// JSON is stored in the requestJson field for custom parsing.
func SetRequestViaLookahead(ctx context.Context, reqRoot *RequestRoot, rootJson []byte) error {
	reqJson := []byte(gjson.GetBytes(rootJson, "request").String())
	var r Request
	rc := RequestCommon{requestJson: reqJson}
	requestType := RequestType(gjson.GetBytes(reqJson, "type").String())
	switch requestType {
	case RequestTypeCanFulfillIntentRequest:
		r = &CanFulfillIntentRequest{RequestCommon: rc}
	case RequestTypeIntentRequest:
		r = &RequestIntentRequest{RequestCommon: rc}
	case RequestTypeSessionEndedRequest:
		r = &RequestSessionEndedRequest{RequestCommon: rc}
	case RequestTypeSessionResumedRequest:
		r = &RequestSessionResumedRequest{RequestCommon: rc}
	case RequestTypeAplUserEvent:
		r = &RequestAplUserEvent{RequestCommon: rc}
	case RequestTypeAlexaAuthorizationGrant:
		r = &RequestAlexaAuthorizationGrant{RequestCommon: rc}
	case RequestTypeAlexaSkillEventSkillAccountLinked:
		r = &RequestAlexaSkillEventSkillAccountLinked{RequestCommon: rc}
	case RequestTypeAlexaSkillEventSkillPermissionAccepted:
		r = &RequestAlexaSkillEventSkillPermissionAccepted{RequestCommon: rc}
	case RequestTypeAlexaSkillEventSkillPermissionChanged:
		r = &RequestAlexaSkillEventSkillPermissionChanged{RequestCommon: rc}
	default:
		r = &rc
	}

	err := json.Unmarshal(reqJson, r)
	if err != nil {
		return err
	}

	reqRoot.Request = r

	return nil
}
