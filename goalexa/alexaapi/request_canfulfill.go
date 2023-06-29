package alexaapi

//
//
// Request: CanFulfillRequest

type CanFulfillIntentRequest struct {
	RequestCommon
	Intent      *Intent     `json:"intent,omitempty"`
}