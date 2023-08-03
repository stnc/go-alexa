package alexaapi

//
//
// Interface: AlexaSkillEvent

//
//
// Request: AlexaSkillEvent.SkillAccountLinked

const RequestTypeAlexaSkillEventSkillAccountLinked = "AlexaSkillEvent.SkillAccountLinked"

type RequestAlexaSkillEventSkillAccountLinked struct {
	RequestCommon
	Body struct {
		AccessToken string `json:"accessToken"`
	} `json:"body"`
}

//
//
// Request: AlexaSkillEvent.SkillEnabled

const RequestTypeAlexaSkillEventSkillEnabled = "AlexaSkillEvent.SkillEnabled"

// no body

//
//
// Request: AlexaSkillEvent.SkillDisabled

const RequestTypeAlexaSkillEventSkillDisabled = "AlexaSkillEvent.SkillDisabled"

// no body

//
//
// Request: AlexaSkillEvent.SkillPermissionAccepted

const RequestTypeAlexaSkillEventSkillPermissionAccepted = "AlexaSkillEvent.SkillPermissionAccepted"

type RequestAlexaSkillEventSkillPermissionAccepted struct {
	RequestCommon
	Body struct {
		AcceptedPermissions []struct {
			Scope string `json:"scope"`
		} `json:"acceptedPermissions"`
		AcceptedPersonPermissions []struct {
			Scope string `json:"scope"`
		} `json:"acceptedPersonPermissions"`
	} `json:"body"`
}

//
//
// Request: AlexaSkillEvent.SkillPermissionChanged

const RequestTypeAlexaSkillEventSkillPermissionChanged = "AlexaSkillEvent.SkillPermissionChanged"

type RequestAlexaSkillEventSkillPermissionChanged struct {
	RequestCommon
	Body struct {
		AcceptedPermissions []struct {
			Scope string `json:"scope"`
		} `json:"acceptedPermissions"`
		AcceptedPersonPermissions []struct {
			Scope string `json:"scope"`
		} `json:"acceptedPersonPermissions"`
	} `json:"body"`
}
