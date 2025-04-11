## Overview
Develop a very simple and very fast alexa application with this go application
Go library to easily handle Alexa custom skill requests, conforming to the official Alexa Skill Kit JSON reference.

A working example for developing [Alexa Skills Kit](http://developer.amazon.com/public/solutions/alexa/alexa-skills-kit)

It contains examples and tests and mocks  [see] (https://github.com/stnc/mygoalexa/tree/master/example)

## Requirements && Installation

1. [Golang](https://golang.org/) v1.6+
2. `go get github.com/stnc/mygoalexa`


## faster example 


```golang
package main

import (
"context"
"encoding/json"
"fmt"
"github.com/stnc/mygoalexa"
"github.com/stnc/mygoalexa/alexaapi"
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
skill := mygoalexa.NewSkill("amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa258929aa")
skill.RegisterHandlers(&LaunchNew{})
http.HandleFunc("/alexa", skill.ServeHTTP)
var port string = "9095"
fmt.Println("server running localhost:" + port)
http.ListenAndServe(":"+port, nil)
}
func (h *LaunchNew) Handle(ctx context.Context, skill *mygoalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

	requestType := requestRoot.Request.GetType()
	fmt.Println(requestType)

	var response alexaapi.ResponseRoot
	var builder Builder

	if requestType == "LaunchRequest" {
		text := "Hi! Welcome to Diet Application"
		title := "diet reminder"
		reprompt := "Welcome Diet Application"
		builder.OutputSpeech(text).Card(title, text).Reprompt(reprompt).EndSession(false)
	}
	if requestType == "IntentRequest" {
		text := "Ok Save successful "
		title := "Diet Reminder"
		reprompt := "You may want to continue the conversation. I am still listening. How can I help you?"
		builder.OutputSpeech(text).Card(title, text).Reprompt(reprompt)
	}

	responseJson, _ := json.Marshal(builder)
	json.Unmarshal([]byte(responseJson), &response)
	return &response, nil
}
func (h *LaunchNew) CanHandle(ctx context.Context, skill *mygoalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
return true
}
```


### Get Intent 

```golang

requestJson := requestRoot.Request.GetRequestJson()

var requestIntent alexaapi.RequestIntentRequest

json.Unmarshal(requestJson, &requestIntent)

numberOfPeopleIntentValue := requestIntent.Intent.Slots["NumberOfPeople"].Value

fmt.Println(numberOfPeopleIntentValue)

intent := GetIntent(requestRoot, "RemindTime")

fmt.Println(intent)

```


## legacy example

```golang

type LaunchReq struct{}

func main() {
skill := mygoalexa.NewSkill("amzn1.ask.skill.d89b3e52-2d85-4693-a664-bcaa25892900")
skill.RegisterHandlers(&LaunchReq{})
http.HandleFunc("/alexa", skill.ServeHTTP)
var port string = "9095"
fmt.Println("server running localhost:" + port)
http.ListenAndServe(":"+port, nil)
}

func (h *LaunchReq) Handle(ctx context.Context, skill *mygoalexa.Skill, requestRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {

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
               fmt.Println(numberOfPeopleIntentValue)

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
func (h *LaunchReq) CanHandle(ctx context.Context, skill *mygoalexa.Skill, requestRoot *alexaapi.RequestRoot) bool {
return true
}
```
## Test

required postman, json and docker also [see here ](https://github.com/stnc/mygoalexa/tree/master/example/tools) for ready postresql application
.

## FORKED

required postman, json and docker also [see here ](https://github.com/stnc/mygoalexa/tree/master/example/tools) for ready postresql application


[FORKED by ](https://github.com/aivahealth/goalexa)

