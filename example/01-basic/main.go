package main

import (
	"context"
	"fmt"
	"github.com/aivahealth/goalexa"
	"github.com/aivahealth/goalexa/alexaapi"
	"github.com/joho/godotenv"
	"net/http"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

type LaunchReq struct{}

func main() {

	skill := goalexa.NewSkill("amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa25892900")
	skill.RegisterHandlers(&LaunchReq{})
	http.HandleFunc("/alexa", skill.ServeHTTP)
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
		myCard.Type = alexaapi.CardTypeSimple
		response.Response.Card = &myCard

		var myReprompt alexaapi.Reprompt
		myReprompt.OutputSpeech = &myOutputSpeech
		response.Response.Reprompt = &myReprompt

		response.Response.ShouldEndSession = &x
	}

	return &response, nil
}
func (h *LaunchReq) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}
