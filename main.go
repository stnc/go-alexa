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

/*
utterance = tr = soyleyis ifade bicimi , ses cikarma  adIrIns speak
slot = yuva , yarik, yerlestirmek sLat

https://www.golangprograms.com/go-language/struct.html  https://kevin-yang.medium.com/golang-embedded-structs-b9d20aadea84
https://www.golangprograms.com/go-language/interface.html -- lesson work ? -> Interface Accepting Address of the Variable == this example

https://github.com/waringer/Alexa-Radio --bundan model skil falan olsuturma ornekleri al

https://github.com/go-alexa/alexa/blob/master/events/events.go#L64

https://github.com/jsgoecke/lambda-go  js json kodlar var

https://github.com/ericdaugherty/alexa-skills-kit-golang/blob/master/alexa_test.go   test yazma ornegi
https://github.com/nraboy/alexa-slick-dealer/blob/master/main_test.go bu da test

https://github.com/go-alexa/alexa/tree/master  burada da test var

https://github.com/go-alexa/alexa/blob/master/server/server_test.go   buradaki function icinde funtion teknigini mutlaka ogren

go ile rouer yazmak https://benhoyt.com/writings/go-routing/ ve https://golangdocs.com/golang-mux-router  konu hakkinda iyisi


//https://www.tutorialspoint.com/golang-program-that-uses-structs-as-map-keys


//https://github.com/NerdCademyDev/golang  burada generic var
//https://gosamples.dev/enum/ https://articles.wesionary.team/working-with-constants-and-iota-in-golang-460c64792d40  const ile auto inc

*/

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

type LaunchReq struct{}

func main() {

	skill := goalexa.NewSkill("amzn1.ask.skill.72d8ce35-6532-481f-aecb-b7149015f763")
	skill.RegisterHandlers(&LaunchReq{})

	http.HandleFunc("/", skill.ServeHTTP)
	var port string = "9091"
	fmt.Println("server running localhost:" + port)

	http.ListenAndServe(":"+port, nil)

}

func (h *LaunchReq) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	intentName := requestRoot.Request.GetType()
	fmt.Println(intentName)

	var response alexaapi.ResponseRoot
	response.Version = "1"
	x := true
	response.SessionAttributes = make(map[string]interface{})
	response.SessionAttributes["read"] = true
	response.SessionAttributes["category"] = true

	requestJson := requestRoot.Request.GetRequestJson()
	trash := map[string]any{}
	json.Unmarshal(requestJson, &trash)
	empJSON, err := json.MarshalIndent(trash, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON))

	if intentName == "LaunchRequest" {
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

	if intentName == "IntentRequest" {
		//if ($EchoReqObj->request->intent->name == "CatFeederFeed"){
		//
		//}

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
