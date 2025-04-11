package alexaapi

//
//
// Interface: Alexa.Authorization

// Request: Alexa.Authorization.Grant

const RequestTypeAlexaAuthorizationGrant = "Alexa.Authorization.Grant"

type RequestAlexaAuthorizationGrant struct {
	RequestCommon
	Body struct {
		Grant struct {
			Code string    `json:"code"`
			Type GrantType `json:"type"`
		} `json:"grant"`
	} `json:"body"`
}

//
//
// Envelope Payload: AcceptGrant

const PayloadNameAcceptGrant PayloadName = "AcceptGrant"

type PayloadAcceptGrant struct {
	Grant struct {
		Type GrantType `json:"type"`
		Code string    `json:"code"`
	} `json:"grant"`
	Grantee struct {
		Type  GranteeType `json:"type"`
		Token string      `json:"token"`
	} `json:"grantee"`
}

type GrantType string

const (
	GrantTypeUnspecified       GrantType = ""
	GrantTypeAuthorizationCode GrantType = "OAuth2.AuthorizationCode"
)

type GranteeType string

const (
	GranteeTypeUnspecified GranteeType = ""
	GranteeTypeBearerToken GranteeType = "BearerToken"
)

//
//
// Envelope Payload: AcceptGrant.Response

// (empty payload)
const PayloadNameAcceptGrantResponse PayloadName = "AcceptGrant.Response"
