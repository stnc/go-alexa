////////////// selman

package main

import (
	cms "avia/app/controller"
	"avia/app/domain/entity"
	repository "avia/app/domain/repository"
	"avia/goalexa"
	"avia/goalexa/alexaapi"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
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

type LaunchNew struct{}

func main() {

	//
	db := repository.DbConnect()
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()

	reminder := cms.InitReminderControl(services.Reminder)

	skill := goalexa.NewSkill("amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa258929aa")
	skill.RegisterHandlers(&LaunchNew{})
	http.HandleFunc("/alexa", skill.ServeHTTP)
	http.HandleFunc("/list", reminder.Index)
	var port string = "9095"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
func (h *LaunchNew) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	requestType := requestRoot.Request.GetType()
	fmt.Println(requestType)

	var response alexaapi.ResponseRoot
	var builder goalexa.Builder

	if requestType == "LaunchRequest" {
		text := "Hi! Welcome to Diet Application"
		title := "diet reminder"
		reprompt := "Welcome Diet Application"
		builder.OutputSpeech(text).Card(title, text).Reprompt(reprompt).EndSession(false)

	}
	if requestType == "IntentRequest" {
		text := "Ok Save successful "
		title := "diet reminder"
		reprompt := "You may want to continue the conversation. I am still listening. How can I help you?"
		builder.OutputSpeech(text).Card(title, text).Reprompt(reprompt)
		// intent read == legacy way
		//requestJson := requestRoot.Request.GetRequestJson()
		//var requestIntent alexaapi.RequestIntentRequest
		//json.Unmarshal(requestJson, &requestIntent)
		//numberOfPeopleIntentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value
		//fmt.Println(numberOfPeopleIntentValue)
		intent := goalexa.GetIntent(requestRoot, "RemindTime")

		var reminder entity.Reminder
		reminder.PersonName = "selman 4"
		reminder.RemindDate = "2023-07-12"
		reminder.RemindTime = "04:00"
		reminder.NumberOfPeople = "4"
		reminder.Email = "dsds@d.com"
		reminder.Phone = "5354543543"
		cms.SaveData(reminder)

		fmt.Println(intent)
	}

	responseJson, _ := json.Marshal(builder)
	json.Unmarshal([]byte(responseJson), &response)
	return &response, nil
}
func (h *LaunchNew) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}
