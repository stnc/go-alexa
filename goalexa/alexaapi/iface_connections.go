package alexaapi

//
//
// Interface: Connections

const (
	DirectiveTypeConnectionsStartConnection DirectiveType = "Connections.StartConnection"
)

// Helper strings to be set as "uri" field in Connections.StartConnection directives
const (
	ConnectionsStartConnectionUriVerifyPerson2 string = "connection://AMAZON.VerifyPerson/2"
)

type ConnectionsStartConnectionInput struct {
	RequestedAuthenticationConfidenceLevel RequestedAuthenticationConfidenceLevel `json:"requestedAuthenticationConfidenceLevel,omitempty"`
}

// level == 400: voice recognition & PIN
// level < 400: voice recognition only
type RequestedAuthenticationConfidenceLevel struct {
	Level        int                                                `json:"level,omitempty"`
	CustomPolicy RequestedAuthenticationConfidenceLevelCustomPolicy `json:"customPolicy,omitempty"`
}

type RequestedAuthenticationConfidenceLevelCustomPolicy struct {
	PolicyName CustomPolicyName `json:"policyName,omitempty"`
}

type CustomPolicyName string

const (
	CustomPolicyNameUnspecified    CustomPolicyName = ""
	CustomPolicyNameVerifyVoicePin CustomPolicyName = "VOICE_PIN"
)

//
//
// Directive: Connections.StartConnection

type DirectiveConnectionsStartConnection struct {
	Type  DirectiveType                    `json:"type"`
	Uri   string                           `json:"uri,omitempty"`
	Input *ConnectionsStartConnectionInput `json:"input,omitempty"`
	Token string                           `json:"token,omitempty"`
}

// This directive is used to verify the user's identity (via PIN)
func CreateDirectiveConnectionsStartConnectionVoicePin(token string) *DirectiveConnectionsStartConnection {
	return &DirectiveConnectionsStartConnection{
		Type:  DirectiveTypeConnectionsStartConnection,
		Uri:   ConnectionsStartConnectionUriVerifyPerson2,
		Token: token,
		Input: &ConnectionsStartConnectionInput{
			RequestedAuthenticationConfidenceLevel: RequestedAuthenticationConfidenceLevel{
				Level: 400,
				CustomPolicy: RequestedAuthenticationConfidenceLevelCustomPolicy{
					PolicyName: CustomPolicyNameVerifyVoicePin,
				},
			},
		},
	}
}
