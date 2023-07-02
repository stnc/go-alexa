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

type HelloWorld struct{}

func main() {

	//var h goalexa.RequestHandler
	//h = Shandler(1)
	//goalexa.HandlerGroup{}

	test := goalexa.NewSkill("amzn1.ask.skill.72d8ce35-6532-481f-aecb-b7149015f763")
	test.RegisterHandlers()
	http.HandleFunc("/", test.ServeHTTP)
	var port string = "8090"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)

	//this struct  not acces -- only access Config value
	//emp1 := goalexa.Skill{
	//	Config:
	//}

	//test.RegisterHandlers()
	//var e1 goalexa.HandlerGroup
	//
	//e1.Handle()

	//var skill goalexa.Skill
	//skill.Config

}

// OnLaunch called with a reqeust is received of type LaunchRequest
func (h *HelloWorld) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	log.Printf("OnLaunch requestId=%s, sessionId=%s", requestRoot.Request.GetRequestId(), skill.Config)
	//speechText := "Welcome to the Alexa Skills Kit, you can say hello"
	//
	//log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)
	//
	//response.SetSimpleCard(cardTitle, speechText)
	//response.SetOutputText(speechText)
	//response.SetRepromptText(speechText)
	//
	//response.ShouldSessionEnd = true
	//
	return nil, nil
}
func (h *HelloWorld) Handle2(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	log.Printf("OnLaunch requestId=%s, sessionId=%s", requestRoot.Request.GetRequestId(), skill.Config)
	//speechText := "Welcome to the Alexa Skills Kit, you can say hello"
	//
	//log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)
	//
	//response.SetSimpleCard(cardTitle, speechText)
	//response.SetOutputText(speechText)
	//response.SetRepromptText(speechText)
	//
	//response.ShouldSessionEnd = true
	//
	return nil, nil
}
func main2() {
	x := true
	//types := alexaapi.CardTypeSimple
	var response alexaapi.ResponseRoot

	response.Version = "1"
	response.SessionAttributes = make(map[string]interface{}) //https://dev.to/rytsh/embed-map-in-json-output-5dnj
	response.SessionAttributes["read"] = true
	response.SessionAttributes["category"] = true

	text := "sddsd"
	types := alexaapi.OutputSpeechTypePlainText

	response.Response.OutputSpeech.Type = types
	response.Response.OutputSpeech.Text = text

	var myCard alexaapi.Card
	myCard.Title = "CatFeeder"
	myCard.Content = "selman content"
	myCard.Type = alexaapi.CardTypeSimple
	response.Response.Card = &myCard

	var myOutputSpeech alexaapi.OutputSpeech
	myOutputSpeech.Text = "ddd"
	myOutputSpeech.Type = alexaapi.OutputSpeechTypePlainText
	response.Response.Reprompt.OutputSpeech = &myOutputSpeech

	response.Response.ShouldEndSession = &x
	//responseJson, _ := json.Marshal(response)
	empJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))
	//_, err = io.Copy(w, bytes.NewReader(responseJson))

}
