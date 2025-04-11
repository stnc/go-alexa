package main

import (
	cms "avia/app/controller"
	"avia/app/domain/entity"
	repository "avia/app/domain/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	 "github.com/stnc/mygoalexa"
	"github.com/stnc/mygoalexa/alexaapi"
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

	skill := mygoalexa.NewSkill("amzn1.ask.skill.27650f59-7c37-441f-8a7d-1a89bf595445")
	skill.RegisterHandlers(&LaunchNew{})

	http.HandleFunc("/alexa", skill.ServeHTTP)
	http.HandleFunc("/list", reminder.Index)

	var port = "9095"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func (h *LaunchNew) Handle(ctx context.Context, skill *mygoalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	requestType := requestRoot.Request.GetType()
	fmt.Println(requestType)

	var response alexaapi.ResponseRoot
	var builder mygoalexa.ResponseRootBuilder

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
		personName := mygoalexa.GetIntent(requestRoot, "PersonName")
		remindDate := mygoalexa.GetIntent(requestRoot, "RemindDate")
		remindTime := mygoalexa.GetIntent(requestRoot, "RemindTime")
		numberOfPeople := mygoalexa.GetIntent(requestRoot, "NumberOfPeople")

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
func (h *LaunchNew) CanHandle(ctx context.Context, skill *mygoalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
	return true
}
