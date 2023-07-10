package main

import (
	"avia/goalexa"
	"avia/goalexa/alexaapi"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

// import 	"github.com/aivahealth/goalexa/alexaapi"

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

type LaunchReq struct{}

func main() {
	skill := goalexa.NewSkill("amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa258929aa")
	skill.RegisterHandlers(&LaunchReq{})
	http.HandleFunc("/alexa/restaurant-bot", skill.ServeHTTP)
	var port string = "9095"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func (h *LaunchReq) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	requestType := requestRoot.Request.GetType()
	fmt.Println(requestType)
	var response alexaapi.ResponseRoot
	response.Version = "7.0"
	x := false

	response.SessionAttributes = make(map[string]interface{})
	response.SessionAttributes["read"] = true
	response.SessionAttributes["category"] = true
	// TODO: how does this code work
	//requestJson := requestRoot.Request.GetRequestJson()

	//requestIntent := map[string]any{}
	//json.Unmarshal(requestJson, &requestIntent)
	//empJSON, err := json.MarshalIndent(requestIntent, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON))

	if requestType == "LaunchRequest" {
		text := "Hi! Welcome to Diet Application"
		var myOutputSpeech alexaapi.OutputSpeech
		myOutputSpeech.Text = text
		myOutputSpeech.Type = alexaapi.OutputSpeechTypeSSML
		myOutputSpeech.SSML = "<speak>" + text + "</speak>"
		response.Response.OutputSpeech = &myOutputSpeech

		var myCard alexaapi.Card
		myCard.Title = "diet reminder "
		myCard.Content = text
		myCard.Type = alexaapi.CardTypeStandard
		response.Response.Card = &myCard
		response.Response.Reprompt.OutputSpeech = &myOutputSpeech
		response.Response.ShouldEndSession = &x
	}

	if requestType == "IntentRequest" {

		//requestJson := requestRoot.Request.GetRequestJson()
		//var requestIntent alexaapi.RequestIntentRequest
		//json.Unmarshal(requestJson, &requestIntent)
		//numberOfPeopleIntentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value
		response.Response.Directives = nil
		text := "Ok Save successful "
		types := alexaapi.OutputSpeechTypePlainText

		var myOutputSpeech alexaapi.OutputSpeech
		myOutputSpeech.Text = text
		myOutputSpeech.Type = types
		response.Response.OutputSpeech = &myOutputSpeech

		var myCard alexaapi.Card
		myCard.Title = "test title"
		myCard.Content = "test context"
		myCard.Type = alexaapi.CardTypeStandard
		response.Response.Card = &myCard
		response.Response.Reprompt.OutputSpeech = &myOutputSpeech
		response.Response.ShouldEndSession = &x
	}
	return &response, nil
}
func (h *LaunchReq) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}

// OutputSpeech will replace any existing text that should be spoken with this new value. If the output
// needs to be constructed in steps or special speech tags need to be used, see the `SSMLTextBuilder`.
func (r *EchoResponse) OutputSpeech(text string) *EchoResponse {
	r.Response.OutputSpeech = &EchoRespPayload{
		Type: "PlainText",
		Text: text,
	}

	return r
}

// Card will add a card to the Alexa app's response with the provided title and content strings.
func (r *EchoResponse) Card(title string, content string) *EchoResponse {
	return r.SimpleCard(title, content)
}

// OutputSpeechSSML will add the text string provided and indicate the speech type is SSML in the response.
// This should only be used if the text to speech string includes special SSML tags.
func (r *EchoResponse) OutputSpeechSSML(text string) *EchoResponse {
	r.Response.OutputSpeech = &EchoRespPayload{
		Type: "SSML",
		SSML: text,
	}

	return r
}

// SimpleCard will indicate that a card should be included in the Alexa companion app as part of the response.
// The card will be shown with the provided title and content.
func (r *EchoResponse) SimpleCard(title string, content string) *EchoResponse {
	r.Response.Card = &EchoRespPayload{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}

	return r
}

// StandardCard will indicate that a card should be shown in the Alexa companion app as part of the response.
// The card shown will include the provided title and content as well as images loaded from the locations provided
// as remote locations.
func (r *EchoResponse) StandardCard(title string, content string, smallImg string, largeImg string) *EchoResponse {
	r.Response.Card = &EchoRespPayload{
		Type:    "Standard",
		Title:   title,
		Content: content,
	}

	if smallImg != "" {
		r.Response.Card.Image.SmallImageURL = smallImg
	}

	if largeImg != "" {
		r.Response.Card.Image.LargeImageURL = largeImg
	}

	return r
}

// LinkAccountCard is used to indicate that account linking still needs to be completed to continue
// using the Alexa skill. This will force an account linking card to be shown in the user's companion app.
func (r *EchoResponse) LinkAccountCard() *EchoResponse {
	r.Response.Card = &EchoRespPayload{
		Type: "LinkAccount",
	}

	return r
}

// Reprompt will send a prompt back to the user, this could be used to request additional information from the user.
func (r *EchoResponse) Reprompt(text string) *EchoResponse {
	r.Response.Reprompt = &EchoReprompt{
		OutputSpeech: EchoRespPayload{
			Type: "PlainText",
			Text: text,
		},
	}

	return r
}

// RepromptSSML is similar to the `Reprompt` method but should be used when the prompt
// to the user should include special speech tags.
func (r *EchoResponse) RepromptSSML(text string) *EchoResponse {
	r.Response.Reprompt = &EchoReprompt{
		OutputSpeech: EchoRespPayload{
			Type: "SSML",
			Text: text,
		},
	}

	return r
}

// EndSession is a convenience method for setting the flag in the response that will
// indicate if the session between the end user's device and the skillserver should be closed.
func (r *EchoResponse) EndSession(flag bool) *EchoResponse {
	r.Response.ShouldEndSession = flag

	return r
}

// RespondToIntent is used to Delegate/Elicit/Confirm a dialog or an entire intent with
// user of alexa. The func takes in name of the dialog, updated intent/intent to confirm
// if any and optional slot value. It prepares a Echo Response to be returned.
// Multiple directives can be returned by calling the method in chain
// (eg. RespondToIntent(...).RespondToIntent(...), each RespondToIntent call appends the
// data to Directives array and will return the same at the end.
func (r *EchoResponse) RespondToIntent(name dialog.Type, intent *EchoIntent, slot *EchoSlot) *EchoResponse {
	directive := EchoDirective{Type: name}
	if intent != nil && name == dialog.ConfirmIntent {
		directive.IntentToConfirm = intent.Name
	} else {
		directive.UpdatedIntent = intent
	}

	if slot != nil {
		if name == dialog.ElicitSlot {
			directive.SlotToElicit = slot.Name
		} else if name == dialog.ConfirmSlot {
			directive.SlotToConfirm = slot.Name
		}
	}
	r.Response.Directives = append(r.Response.Directives, &directive)
	return r
}

func (r *EchoResponse) String() ([]byte, error) {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return jsonStr, nil
}
