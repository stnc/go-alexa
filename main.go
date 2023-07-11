////////////// selman

package main

import (
	"avia/goalexa"
	"avia/goalexa/alexaapi"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
)

// import 	"github.com/aivahealth/goalexa/alexaapi"
//https://github.com/nraboy/alexa-slick-dealer/blob/master/main.go

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

func swap2(px *alexaapi.ResponseRoot, py *alexaapi.ResponseRoot) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}

func swap(px, py *int) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}
func main3() {
	x := int(1)
	y := int(2)
	fmt.Println("x was", x)
	fmt.Println("y was", y)
	swap(&x, &y)

	fmt.Println("x is now", x)
	fmt.Println("y is now", y)

	var res alexaapi.ResponseRoot
	res.Version = "6"

	var response goalexa.Builder
	response.OutputSpeech("dsds")
	response.Version = "6"

	empJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON))

	//swap2(&res, &response)
}

type LaunchReq struct{}
type LaunchNew struct{}

func main() {
	//text := "Hi! Welcome to Diet Application2"
	//title2 := "diet reminder2"
	//var response goalexa.Builder
	//response.OutputSpeech(text).SimpleCard(title2, text)
	//empJSON, err := json.MarshalIndent(response, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON))
	skill := goalexa.NewSkill("amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa258929aa")
	skill.RegisterHandlers(&LaunchNew{})
	http.HandleFunc("/alexi", skill.ServeHTTP)
	var port string = "9095"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
func (h *LaunchNew) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	requestType := requestRoot.Request.GetType()
	fmt.Println(requestType)

	var response alexaapi.ResponseRoot
	var response2 goalexa.Builder
	x := false

	if requestType == "LaunchRequest" {
		//text := "Hi! Welcome to Diet Application2"
		//title2 := "diet reminder2"
		//response2.OutputSpeech(text).Card(title2, text).Reprompt(text)
		text := "Ok Save successful "
		title2 := "diet reminder2"
		reprompt := "You may want to continue the conversation. I am still listening. How can I help you?"
		response2.OutputSpeech(text).Card(title2, text).Reprompt(reprompt)
		//requestJson := requestRoot.Request.GetRequestJson()
		//var requestIntent alexaapi.RequestIntentRequest
		//json.Unmarshal(requestJson, &requestIntent)
		//numberOfPeopleIntentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value
		response.Response.Directives = nil
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

		var myReprompt alexaapi.Reprompt
		myReprompt.OutputSpeech = &myOutputSpeech
		response.Response.Reprompt = &myReprompt

		response.Response.ShouldEndSession = &x
	}
	//response2 = &response
	if requestType == "IntentRequest" {
		text := "Ok Save successful "
		title2 := "diet reminder2"
		reprompt := "You may want to continue the conversation. I am still listening. How can I help you?"
		response2.OutputSpeech(text).Card(title2, text).Reprompt(reprompt)
		//requestJson := requestRoot.Request.GetRequestJson()
		//var requestIntent alexaapi.RequestIntentRequest
		//json.Unmarshal(requestJson, &requestIntent)
		//numberOfPeopleIntentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value
		response.Response.Directives = nil
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

		var myReprompt alexaapi.Reprompt
		myReprompt.OutputSpeech = &myOutputSpeech
		response.Response.Reprompt = &myReprompt

		response.Response.ShouldEndSession = &x
	}

	//empJSON, err := json.MarshalIndent(response, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON))
	//
	//empJSON2, err := json.MarshalIndent(response2, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON2))

	responseJson, err := json.Marshal(response)
	if err != nil {
		fmt.Println("ServeHTTP failed", zap.Error(err))
		return nil, nil
	}

	fmt.Println(string(responseJson))

	return &response, nil
}
func (h *LaunchNew) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}

//func (h *LaunchReq) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {
//
//	requestType := requestRoot.Request.GetType()
//	fmt.Println(requestType)
//	var response alexaapi.ResponseRoot
//	response.Version = "7.0"
//	x := false
//
//	response.SessionAttributes = make(map[string]interface{})
//	response.SessionAttributes["read"] = true
//	response.SessionAttributes["category"] = true
//
//	// TODO: how does this code work
//	//requestJson := requestRoot.Request.GetRequestJson()
//	//requestIntent := map[string]any{}
//	//json.Unmarshal(requestJson, &requestIntent)
//	//empJSON, err := json.MarshalIndent(requestIntent, "", "  ")
//	//if err != nil {
//	//	log.Fatalf(err.Error())
//	//}
//	//fmt.Printf("MarshalIndent funnction only intent output\n %s\n", string(empJSON))
//
//	if requestType == "LaunchRequest" {
//		text := "Hi! Welcome to Diet Application"
//
//		var myOutputSpeech alexaapi.OutputSpeech
//		myOutputSpeech.Text = text
//		myOutputSpeech.Type = alexaapi.OutputSpeechTypeSSML
//		myOutputSpeech.SSML = "<speak>" + text + "</speak>"
//		response.Response.OutputSpeech = &myOutputSpeech
//
//		var myCard alexaapi.Card
//		myCard.Title = "diet reminder "
//		myCard.Content = text
//		myCard.Type = alexaapi.CardTypeStandard
//		response.Response.Card = &myCard
//
//		var myReprompt alexaapi.Reprompt
//		myReprompt.OutputSpeech = &myOutputSpeech
//		response.Response.Reprompt = &myReprompt
//		response.Response.ShouldEndSession = &x
//
//	}
//
//	if requestType == "IntentRequest" {
//
//		//requestJson := requestRoot.Request.GetRequestJson()
//		//var requestIntent alexaapi.RequestIntentRequest
//		//json.Unmarshal(requestJson, &requestIntent)
//		//numberOfPeopleIntentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value
//		response.Response.Directives = nil
//		text := "Ok Save successful "
//		types := alexaapi.OutputSpeechTypePlainText
//
//		var myOutputSpeech alexaapi.OutputSpeech
//		myOutputSpeech.Text = text
//		myOutputSpeech.Type = types
//		response.Response.OutputSpeech = &myOutputSpeech
//
//		var myCard alexaapi.Card
//		myCard.Title = "test title"
//		myCard.Content = "test context"
//		myCard.Type = alexaapi.CardTypeStandard
//		response.Response.Card = &myCard
//
//		var myReprompt alexaapi.Reprompt
//		myReprompt.OutputSpeech = &myOutputSpeech
//		response.Response.Reprompt = &myReprompt
//
//		response.Response.ShouldEndSession = &x
//	}
//	return &response, nil
//}
//func (h *LaunchReq) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
//	return true
//}
