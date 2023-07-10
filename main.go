package main

import (
	"avia/goalexa"
	"avia/goalexa/alexaapi"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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


test yazmak
https://github.com/jakubsuchy/amazon-alexa-php

go kucuk kodlamar ornegi
*/

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

		text := "Hi! Welcome to sesame Restaurant for selmantunc.com How can I help you today? amazon"

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
		//TODO remain works ... save data and helper function, look go alaexa skill project
		requestJson := requestRoot.Request.GetRequestJson()
		var requestIntent alexaapi.RequestIntentRequest
		json.Unmarshal(requestJson, &requestIntent)
		fmt.Println(requestIntent.Intent.Name)
		intentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value
		fmt.Println(intentValue)

		response.Response.Directives = nil
		text := "Ok Save successful "
		types := alexaapi.OutputSpeechTypePlainText

		var myOutputSpeech alexaapi.OutputSpeech

		myOutputSpeech.Text = text
		myOutputSpeech.Type = types

		response.Response.OutputSpeech = &myOutputSpeech

		var myCard alexaapi.Card
		myCard.Title = "test title"
		myCard.Content = "bal content"
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

func fileAdd(data []string) {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range data {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
