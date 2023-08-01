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

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

type LaunchNew struct{}

func main() {
	db := repository.DbConnect()
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()
	reminder := cms.InitReminderControl(services.Reminder)

	skill := goalexa.NewSkill("amzn1.ask.skill.27650f59-7c37-441f-8a7d-1a89bf595445")
	skill.RegisterHandlers(&LaunchNew{})

	http.HandleFunc("/alexa", skill.ServeHTTP)
	http.HandleFunc("/list", reminder.Index)

	var port = "9095"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func (h *LaunchNew) Handle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	requestType := requestRoot.Request.GetType()
	fmt.Println(requestType)

	var response alexaapi.ResponseRoot
	var builder goalexa.ResponseRootBuilder

	if requestType == "LaunchRequest" {
		text := "Hi! Welcome to Diet Application"
		title := "diet reminder"
		reprompt := "Welcome Diet Application"
		builder.OutputSpeech(text).Card(title, text).Reprompt(reprompt).AddAudioPlayer("", "", 1).EndSession(false)
	}

	if requestType == "IntentRequest" {
		text := "Ok Save successful "
		title := "Diet Reminder"
		reprompt := "You may want to continue the conversation. I am still listening. How can I help you?"
		builder.OutputSpeech(text).Card(title, text).Reprompt(reprompt)
		personName := goalexa.GetIntent(requestRoot, "PersonName")
		remindDate := goalexa.GetIntent(requestRoot, "RemindDate")
		remindTime := goalexa.GetIntent(requestRoot, "RemindTime")
		numberOfPeople := goalexa.GetIntent(requestRoot, "NumberOfPeople")

		var reminder entity.Reminder
		reminder.PersonName = personName
		reminder.RemindDate = remindDate
		reminder.RemindTime = remindTime
		reminder.NumberOfPeople = numberOfPeople
		cms.SaveData(reminder)
	}

	responseJson, _ := json.Marshal(builder)
	json.Unmarshal([]byte(responseJson), &response)
	return &response, nil
}
func (h *LaunchNew) CanHandle(ctx context.Context, skill *goalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}
