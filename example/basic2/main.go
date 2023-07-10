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
	"time"
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

		requestJson := requestRoot.Request.GetRequestJson()
		var requestIntent alexaapi.RequestIntentRequest
		json.Unmarshal(requestJson, &requestIntent)

		numberOfPeopleIntentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value
		reservationDateIntentValue := requestIntent.Intent.Slots["RemindDate"].Value
		reservationTimeIntentValue := requestIntent.Intent.Slots["RemindTime"].Value
		personNameIntentValue := requestIntent.Intent.Slots["PersonName"].Value
		currentTime := time.Now()
		today := currentTime.Format("2006.01.02 15:04:05")
		data := []string{
			"<-----------------Date " + today + "------------------------------>",
			"Name: " + personNameIntentValue + " ",
			"Remind Date : " + reservationDateIntentValue + " ",
			"Remind Time : " + reservationTimeIntentValue + " ",
			"How Many People : " + numberOfPeopleIntentValue + " ",
			"<------------------------- THE END -------------------------------------->",
			"                                                                 ",
		}
		dataAdd(data)

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

func dataAdd(data []string) {
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
