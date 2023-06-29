package alexaapi

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/tidwall/gjson"
)

/*
Alexa may send a Directive message to the skill service
(Note: this kind of Directive is distinct from Custom Skill directives sent by the skill service to Alexa.)

The skill service may send an Event message to Alexa

Both Directive and Event messages are sent in the same format
*/

// Amazon docs don't name this, but it's used by both Directive and Event
type Envelope struct {
	Header EnvelopeHeader `json:"header"`

	// If a better type isn't found, map[string]any will be used
	Payload any `json:"payload"`

	// The original payload is stored here in case custom parsing is needed
	PayloadJson []byte `json:"-"`
}

type PayloadName string

const PayloadVersion3 = "3"

type EnvelopeHeader struct {
	Namespace      Interface   `json:"namespace"`
	Name           PayloadName `json:"name"`
	MessageId      string      `json:"messageId"`
	PayloadVersion string      `json:"payloadVersion"`
}

func SetEnvelopePayloadViaLookahead(ctx context.Context, envelope *Envelope, envelopeJson []byte) error {
	if envelope == nil {
		return errors.New("envelope is nil")
	}

	payloadJson := []byte(gjson.GetBytes(envelopeJson, "payload").String())

	switch envelope.Header.Namespace {
	case InterfaceAlexaAuthorization:
		switch envelope.Header.Name {
		case PayloadNameAcceptGrant:
			var payload PayloadAcceptGrant
			err := json.Unmarshal(payloadJson, &payload)
			if err != nil {
				return err
			}
			envelope.Payload = &payload
		}
	}

	envelope.Payload = map[string]any{}
	err := json.Unmarshal(payloadJson, &envelope.Payload)
	if err != nil {
		return err
	}

	envelope.PayloadJson = payloadJson

	return nil
}
