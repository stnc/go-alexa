package alexaapi

//
//
// Request: SessionResumedRequest

type RequestSessionResumedRequest struct {
	RequestCommon
	Cause *SessionResumedRequestCause `json:"cause,omitempty"`
}

// Polymorphic, make sure to consider Type
type SessionResumedRequestCause struct {
	Type SessionResumedRequestCauseType `json:"type"`

	// ConnectionCompleted
	Token  string                       `json:"token,omitempty"`
	Status *SessionResumedRequestStatus `json:"status,omitempty"`
	Result *SessionResumedRequestResult `json:"result,omitempty"`
}

type SessionResumedRequestCauseType string

const (
	SessionResumedRequestCauseTypeUnspecified         SessionResumedRequestCauseType = ""
	SessionResumedRequestCauseTypeConnectionCompleted SessionResumedRequestCauseType = "ConnectionCompleted"
)

type SessionResumedRequestStatusCode string

const (
	SessionResumedRequestStatusCodeUnspecified         SessionResumedRequestStatusCode = ""
	SessionResumedRequestStatusCodeOK                  SessionResumedRequestStatusCode = "200"
	SessionResumedRequestStatusCodeBadRequest          SessionResumedRequestStatusCode = "400"
	SessionResumedRequestStatusCodeForbidden           SessionResumedRequestStatusCode = "403"
	SessionResumedRequestStatusCodeNotFound            SessionResumedRequestStatusCode = "404"
	SessionResumedRequestStatusCodeInternalServerError SessionResumedRequestStatusCode = "500"
)

type SessionResumedRequestStatus struct {
	Code    SessionResumedRequestStatusCode `json:"code"`
	Message string                          `json:"message"`
}

type SessionResumedRequestResultStatus string

const (
	SessionResumedRequestResultStatusUnspecified SessionResumedRequestResultStatus = ""
	SessionResumedRequestResultStatusAchieved    SessionResumedRequestResultStatus = "ACHIEVED"
	SessionResumedRequestResultStatusNotAchieved SessionResumedRequestResultStatus = "NOT_ACHIEVED"
	SessionResumedRequestResultStatusNotEnabled  SessionResumedRequestResultStatus = "NOT_ENABLED"
)

type SessionResumedRequestResultReason string

const (
	SessionResumedRequestResultReasonUnspecified                SessionResumedRequestResultReason = ""
	SessionResumedRequestResultReasonMethodLockout              SessionResumedRequestResultReason = "METHOD_LOCKOUT"
	SessionResumedRequestResultReasonVerificationMethodNotSetup SessionResumedRequestResultReason = "VERIFICATION_METHOD_NOT_SETUP"
	SessionResumedRequestResultReasonNotMatch                   SessionResumedRequestResultReason = "NOT_MATCH"
)

type SessionResumedRequestResult struct {
	Status SessionResumedRequestResultStatus `json:"status"`
	Reason SessionResumedRequestResultReason `json:"reason"`
}
