package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// //example 1

type Respond struct {
	EchoResponse
}
type EchoResponse struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          EchoRespBody           `json:"response"`
}

// EchoRespBody contains the body of the response to be sent back to the Alexa service.
// This includes things like the text that should be spoken or any cards that should
// be shown in the Alexa companion app.

type EchoRespBody struct {
	OutputSpeech     *EchoRespPayload `json:"outputSpeech,omitempty"`
	Card             *EchoRespPayload `json:"card,omitempty"`
	ShouldEndSession bool             `json:"shouldEndSession"`
}

// EchoRespPayload contains the interesting parts of the Echo response including text to be spoken,
// card attributes, and images.

type EchoRespPayload struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Text    string `json:"text,omitempty"`
	SSML    string `json:"ssml,omitempty"`
	Content string `json:"content,omitempty"`
}

// OutputSpeech will replace any existing text that should be spoken with this new value. If the output
// needs to be constructed in steps or special speech tags need to be used, see the `SSMLTextBuilder`.

func (r *Respond) OutputSpeech(text string) *Respond {
	r.Response.OutputSpeech = &EchoRespPayload{
		Type: "PlainText",
		Text: text,
	}

	return r
}

// Card will add a card to the Alexa app's response with the provided title and content strings.

func (r *Respond) Card(title string, content string) *Respond {
	return r.SimpleCard(title, content)
}

// SimpleCard will indicate that a card should be included in the Alexa companion app as part of the response.
// The card will be shown with the provided title and content.

func (r *Respond) SimpleCard(title string, content string) *Respond {
	r.Response.Card = &EchoRespPayload{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}

	return r
}

func main() {
	/*
		In methods, the input struct and output value must be the same, otherwise it will not work, the above is an example of this
	*/
	var echoResp Respond
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
	empJSON, err := json.MarshalIndent(echoResp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

	/// Example 2 -> not spawning
	var resp Builder
	resp.OutputSpeechSSML("hello")

}

// //example 2

type ResponseRoot struct {
	Version           string         `json:"version,omitempty"`
	SessionAttributes map[string]any `json:"sessionAttributes,omitempty"`
	Response          Response       `json:"response,omitempty"`
}

type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         Reprompt      `json:"reprompt,omitempty"` //TUNC
	ShouldEndSession *bool         `json:"shouldEndSession,omitempty"`
	Directives       []any         `json:"directives,omitempty"`
}
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}
type CardImage struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}
type Card struct {
	Type    string    `json:"type"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Text    string    `json:"text,omitempty"`
	Image   CardImage `json:"image,omitempty"`
}
type OutputSpeech struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

type Builder struct {
	ResponseRoot
}

// OutputSpeech will replace any existing text that should be spoken with this new value. If the output
// needs to be constructed in steps or special speech tags need to be used, see the `SSMLTextBuilder`.
func (build *Builder) OutputSpeech(text string) *ResponseRoot {
	var sml string = "<speak>" + text + "</speak>"

	var ResponseRoot ResponseRoot
	var myOutputSpeech OutputSpeech
	myOutputSpeech.Text = text
	myOutputSpeech.Type = "SSML"
	myOutputSpeech.SSML = sml
	ResponseRoot.Response.OutputSpeech = &myOutputSpeech
	return &ResponseRoot
}

// SimpleCard will indicate that a card should be included in the Alexa companion app as part of the response.
// The card will be shown with the provided title and content.
func (build *Builder) SimpleCard(title string, content string) *ResponseRoot {
	var myCard Card
	myCard.Title = title
	myCard.Content = content
	myCard.Type = "Simple"
	build.Response.Card = &myCard
	return &build.ResponseRoot
}

// Card will add a card to the Alexa app's response with the provided title and content strings.
func (build *Builder) Card(title string, content string) *ResponseRoot {
	return build.SimpleCard(title, content)
}

// OutputSpeechSSML will add the text string provided and indicate the speech type is SSML in the response.
// This should only be used if the text to speech string includes special SSML tags.
func (build *Builder) OutputSpeechSSML(text string) *ResponseRoot {
	var sml string = "<speak>" + text + "</speak>"
	var myOutputSpeech OutputSpeech
	myOutputSpeech.Type = "SSML"
	myOutputSpeech.SSML = sml
	build.Response.OutputSpeech = &myOutputSpeech
	return &build.ResponseRoot
}

// StandardCard will indicate that a card should be shown in the Alexa companion app as part of the response.
// The card shown will include the provided title and content as well as images loaded from the locations provided
// as remote locations.
func (build *Builder) StandardCard(title string, content string, smallImg string, largeImg string) *ResponseRoot {
	var myCard Card
	myCard.Title = title
	myCard.Content = content
	myCard.Type = "Standard"
	build.Response.Card = &myCard

	if smallImg != "" {
		build.ResponseRoot.Response.Card.Image.SmallImageURL = smallImg
	}

	if largeImg != "" {
		build.ResponseRoot.Response.Card.Image.LargeImageURL = largeImg
	}

	return &build.ResponseRoot
}
