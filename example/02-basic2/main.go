package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aivahealth/goalexa"
	"github.com/aivahealth/goalexa/alexaapi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
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

func dataAdd(data []string) {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
