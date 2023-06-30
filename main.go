package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"avia/goalexa"
	"avia/goalexa/alexaapi"
	"github.com/joho/godotenv"
)

// import 	"github.com/aivahealth/goalexa/alexaapi"

/*
utterance = tr = soyleyis ifade bicimi , ses cikarma  adIrIns speak
slot = yuva , yarik, yerlestirmek sLat

https://www.golangprograms.com/go-language/struct.html
https://www.golangprograms.com/go-language/interface.html -- lesson work ? -> Interface Accepting Address of the Variable == this example


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
	test := goalexa.NewSkill("amzn1.ask.skill.8271ee57-716d-46db-bf6d-684e27ca4052")
	http.HandleFunc("/check", test.ServeHTTP)
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
