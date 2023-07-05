package main

import (
	"avia/goalexa"
	"avia/goalexa/alexaapi"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

type LaunchReq struct{}
type MakeReservationReq struct{}

func main() {

	test := goalexa.NewSkill("amzn1.ask.skill.72d8ce35-6532-481f-aecb-b7149015f763")
	test.RegisterHandlers(&LaunchReq{})

	http.HandleFunc("/", test.ServeHTTP)
	var port string = "9092"
	fmt.Println("server running localhost:" + port)

	http.ListenAndServe(":"+port, nil)

}

func (h *LaunchReq) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {
	x := true
	var response alexaapi.ResponseRoot
	response.Version = "1"
	response.SessionAttributes = make(map[string]interface{})
	response.SessionAttributes["read"] = true
	response.SessionAttributes["category"] = true
	res := requestRoot.Request.GetType()
	fmt.Println(res)

	if res == "LaunchRequest" {
		text := "Welcome to the drug reminder application"
		types := alexaapi.OutputSpeechTypePlainText

		var myOutputSpeech alexaapi.OutputSpeech
		myOutputSpeech.Text = text
		myOutputSpeech.Type = types
		response.Response.OutputSpeech = &myOutputSpeech

		var myCard alexaapi.Card
		myCard.Title = "Drug reminder "
		myCard.Content = text
		myCard.Type = alexaapi.CardTypeSimple
		response.Response.Card = &myCard

		response.Response.Reprompt.OutputSpeech = &myOutputSpeech
		response.Response.ShouldEndSession = &x
	}
	return &response, nil

}
func (h *LaunchReq) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}

// Handle not used, only example
func (h *MakeReservationReq) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {
	var response alexaapi.ResponseRoot
	return &response, nil
}

func (h *MakeReservationReq) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}
