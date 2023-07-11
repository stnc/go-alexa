package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Builder struct {
	ResponseRoot
}

type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	ShouldEndSession *bool         `json:"shouldEndSession,omitempty"`
	Directives       []any         `json:"directives,omitempty"`
}

type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

type Card struct {
	Type    string    `json:"type"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Text    string    `json:"text,omitempty"`
	Image   CardImage `json:"image,omitempty"`
}

type CardImage struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type OutputSpeech struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

type ResponseRoot struct {
	Version           string         `json:"version,omitempty"`
	SessionAttributes map[string]any `json:"sessionAttributes,omitempty"`
	Response          Response       `json:"response,omitempty"`
}

// OutputSpeech will replace any existing text that should be spoken with this new value. If the output
// needs to be constructed in steps or special speech tags need to be used, see the `SSMLTextBuilder`.
func (build *Builder) OutputSpeech(text string) *Builder {
	var sml string = ("<speak>" + text + "</speak>")
	//myOutputSpeech := OutputSpeech{Text: "Alice", SSML: sml, Type: OutputSpeechTypeSSML}
	//b.Response.OutputSpeech = &myOutputSpeech
	//return &b.ResponseRoot

	var myOutputSpeech OutputSpeech
	myOutputSpeech.Text = text
	myOutputSpeech.Type = "SSML"
	myOutputSpeech.SSML = sml
	build.Response.OutputSpeech = &myOutputSpeech
	return build
}

// SimpleCard will indicate that a card should be included in the Alexa companion app as part of the response.
// The card will be shown with the provided title and content.
func (build *Builder) SimpleCard(title string, content string) *Builder {
	var myCard Card
	myCard.Title = title
	myCard.Content = content
	myCard.Type = "Simple"
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
	var sml string = ("<speak>" + text + "</speak>")
	var myOutputSpeech OutputSpeech
	myOutputSpeech.Type = "SSML"
	myOutputSpeech.SSML = sml
	build.Response.OutputSpeech = &myOutputSpeech
	return build
}

// StandardCard will indicate that a card should be shown in the Alexa companion app as part of the response.
// The card shown will include the provided title and content as well as images loaded from the locations provided
// as remote locations.
func (build *Builder) StandardCard(title string, content string, smallImg string, largeImg string) *Builder {
	var myCard Card
	myCard.Title = title
	myCard.Content = content
	myCard.Type = "standart"
	build.Response.Card = &myCard

	if smallImg != "" {
		build.ResponseRoot.Response.Card.Image.SmallImageURL = smallImg
	}

	if largeImg != "" {
		build.ResponseRoot.Response.Card.Image.LargeImageURL = largeImg
	}

	return build
}

// LinkAccountCard is used to indicate that account linking still needs to be completed to continue
// using the Alexa skill. This will force an account linking card to be shown in the user's companion app.
func (build *Builder) LinkAccountCard() *Builder {
	var myCard Card
	myCard.Type = "LinkAccount"
	build.Response.Card = &myCard
	return build
}

// Reprompt will send a prompt back to the user, this could be used to request additional information from the user.
func (build *Builder) Reprompt(text string) *Builder {
	var myOutputSpeech OutputSpeech
	myOutputSpeech.Text = text
	myOutputSpeech.Type = "PlainText"

	var myReprompt Reprompt
	myReprompt.OutputSpeech = &myOutputSpeech
	build.Response.Reprompt = &myReprompt

	return build
}

// RepromptSSML is similar to the `Reprompt` method but should be used when the prompt
// to the user should include special speech tags.
func (build *Builder) RepromptSSML(text string) *Builder {
	var myOutputSpeech OutputSpeech
	myOutputSpeech.Text = text
	myOutputSpeech.Type = "SSML"
	build.Response.Reprompt.OutputSpeech = &myOutputSpeech
	return build
}

// EndSession is a convenience method for setting the flag in the response that will
// indicate if the session between the end user's device and the skill should be closed.
func (build *Builder) EndSession(flag bool) *Builder {
	build.Response.ShouldEndSession = &flag
	return build
}
func main() {
	text := "Hi! Welcome to Diet Application2"
	title2 := "diet reminder2"
	var response Builder
	response.OutputSpeech(text).SimpleCard(title2, text)
	empJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON))
}
