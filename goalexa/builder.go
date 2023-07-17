package goalexa

import (
	"encoding/json"
	"github.com/aivahealth/goalexa/alexaapi"
)

type Builder alexaapi.ResponseRoot

// OutputSpeech will replace any existing text that should be spoken with this new value. If the output
// needs to be constructed in steps or special speech tags need to be used, see the `SSMLTextBuilder`.
func (build *Builder) OutputSpeech(text string) *Builder {
	var sml string = "<speak>" + text + "</speak>"
	//myOutputSpeech2 := alexaapi.OutputSpeech{Text: "Alice", SSML: sml, Type: alexaapi.OutputSpeechTypeSSML}
	//build.Response.OutputSpeech = &myOutputSpeech2
	//return build
	var myOutputSpeech alexaapi.OutputSpeech
	myOutputSpeech.Text = text
	myOutputSpeech.Type = alexaapi.OutputSpeechTypeSSML
	myOutputSpeech.SSML = sml
	build.Response.OutputSpeech = &myOutputSpeech
	return build
}

// SimpleCard will indicate that a card should be included in the Alexa companion app as part of the response.
// The card will be shown with the provided title and content.
func (build *Builder) SimpleCard(title string, content string) *Builder {
	var myCard alexaapi.Card
	myCard.Title = title
	myCard.Content = content
	myCard.Type = alexaapi.CardTypeSimple
	build.Response.Card = &myCard
	return build
}

// Card will add a card to the Alexa app's response with the provided title and content strings.
func (build *Builder) Card(title string, content string) *Builder {
	return build.SimpleCard(title, content)
}

// OutputSpeechSSML will add the text string provided and indicate the speech type is SSML in the response.
// This should only be used if the text to speech string includes special SSML tags.
func (build *Builder) OutputSpeechSSML(text string) *Builder {
	//var sml string = EscapeSSMLText("<speak>" + text + "</speak>")
	var sml string = "<speak>" + text + "</speak>"
	var myOutputSpeech alexaapi.OutputSpeech
	myOutputSpeech.Type = alexaapi.OutputSpeechTypeSSML
	myOutputSpeech.SSML = sml
	build.Response.OutputSpeech = &myOutputSpeech
	return build
}

// StandardCard will indicate that a card should be shown in the Alexa companion app as part of the response.
// The card shown will include the provided title and content as well as images loaded from the locations provided
// as remote locations.
func (build *Builder) StandardCard(title string, content string, smallImg string, largeImg string) *Builder {
	var myCard alexaapi.Card
	myCard.Title = title
	myCard.Content = content
	myCard.Type = alexaapi.CardTypeStandard
	build.Response.Card = &myCard

	if smallImg != "" {
		build.Response.Card.Image.SmallImageURL = smallImg
	}

	if largeImg != "" {
		build.Response.Card.Image.LargeImageURL = largeImg
	}

	return build
}

// LinkAccountCard is used to indicate that account linking still needs to be completed to continue
// using the Alexa skill. This will force an account linking card to be shown in the user's companion app.
func (build *Builder) LinkAccountCard() *Builder {
	var myCard alexaapi.Card
	myCard.Type = alexaapi.CardTypeLinkAccount
	build.Response.Card = &myCard
	return build
}

// Reprompt will send a prompt back to the user, this could be used to request additional information from the user.
func (build *Builder) Reprompt(text string) *Builder {
	var myOutputSpeech alexaapi.OutputSpeech
	var myReprompt alexaapi.Reprompt

	myOutputSpeech.Text = text
	myOutputSpeech.Type = alexaapi.OutputSpeechTypePlainText

	myReprompt.OutputSpeech = &myOutputSpeech
	build.Response.Reprompt = &myReprompt

	return build
}

// RepromptSSML is similar to the `Reprompt` method but should be used when the prompt
// to the user should include special speech tags.
func (build *Builder) RepromptSSML(text string) *Builder {
	var myOutputSpeech alexaapi.OutputSpeech
	myOutputSpeech.Text = text
	myOutputSpeech.Type = alexaapi.OutputSpeechTypeSSML
	build.Response.Reprompt.OutputSpeech = &myOutputSpeech
	return build
}

// EndSession is a convenience method for setting the flag in the response that will
// indicate if the session between the end user's device and the skill should be closed.
func (build *Builder) EndSession(flag bool) *Builder {
	build.Response.ShouldEndSession = &flag
	return build
}

// GetIntent to quickly reach the intent value from the response
func GetIntent(requestRoot *alexaapi.RequestRoot, name string) string {
	requestJson := requestRoot.Request.GetRequestJson()
	var requestIntent alexaapi.RequestIntentRequest
	json.Unmarshal(requestJson, &requestIntent)
	return requestIntent.Intent.Slots[name].Value
}
