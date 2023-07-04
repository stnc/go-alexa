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
type MakeReservationReq struct{}

func main() {

	//var h goalexa.RequestHandler
	//h = Shandler(1)
	//goalexa.HandlerGroup{}

	test := goalexa.NewSkill("amzn1.ask.skill.72d8ce35-6532-481f-aecb-b7149015f763")
	test.RegisterHandlers(&LaunchReq{}, &MakeReservationReq{})

	http.HandleFunc("/", test.ServeHTTP)
	var port string = "9092"
	fmt.Println("server running localhost:" + port)

	http.ListenAndServe(":"+port, nil)

}

func (h *LaunchReq) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	x := false
	var response alexaapi.ResponseRoot

	//response.Response.CanFulfillIntent.Slots

	response.Version = "1"
	response.SessionAttributes = make(map[string]interface{})
	response.SessionAttributes["read"] = true
	response.SessionAttributes["category"] = true

	text := "Hi man How are you "
	types := alexaapi.OutputSpeechTypePlainText

	var myOutputSpeech alexaapi.OutputSpeech

	myOutputSpeech.Text = text
	myOutputSpeech.Type = types

	response.Response.OutputSpeech = &myOutputSpeech

	var myCard alexaapi.Card
	myCard.Title = "deneme"
	myCard.Content = text
	myCard.Type = alexaapi.CardTypeSimple

	response.Response.Card = &myCard
	response.Response.Reprompt.OutputSpeech = &myOutputSpeech

	response.Response.ShouldEndSession = &x

	return &response, nil

}
func (h *LaunchReq) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}
func (h *MakeReservationReq) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	res := requestRoot.Request.GetType()
	if res == "IntentRequest" {
		x := true
		var response alexaapi.ResponseRoot
		response.Version = "1"
		response.SessionAttributes = make(map[string]interface{})
		response.SessionAttributes["read"] = true
		response.SessionAttributes["category"] = true

		text := "farkli bir  "
		types := alexaapi.OutputSpeechTypePlainText

		var myOutputSpeech alexaapi.OutputSpeech

		myOutputSpeech.Text = text
		myOutputSpeech.Type = types

		response.Response.OutputSpeech = &myOutputSpeech

		var myCard alexaapi.Card
		myCard.Title = "CatFeeder"
		myCard.Content = "selman content"
		myCard.Type = alexaapi.CardTypeSimple
		response.Response.Card = &myCard

		response.Response.Reprompt.OutputSpeech = &myOutputSpeech

		response.Response.ShouldEndSession = &x

		//empJSON, err := json.MarshalIndent(reqRoot, "", "  ")
		//if err != nil {
		//	log.Fatalf(err.Error())
		//}
		//fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

		return &response, nil
	}

}
func (h *MakeReservationReq) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}

//empJSON, err := json.MarshalIndent(reqRoot, "", "  ")
//if err != nil {
//	log.Fatalf(err.Error())
//}
//fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))
