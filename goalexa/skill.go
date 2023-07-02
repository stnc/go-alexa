package goalexa

import (
	"avia/goalexa/alexaapi"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RequestHandler interface {
	CanHandle(context.Context, *Skill, *alexaapi.RequestRoot) bool
	Handle(context.Context, *Skill, *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error)
}

type HandlerGroup []RequestHandler

func (hg HandlerGroup) Handle(ctx context.Context, s *Skill, reqRoot *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {
	for _, h := range hg {
		if h.CanHandle(ctx, s, reqRoot) {
			return h.Handle(ctx, s, reqRoot)
		}
	}

	return nil, fmt.Errorf("No handler found for request (%q)", reqRoot.Request.GetType())
}

type Skill struct {
	Config any

	applicationId string
	handlers      HandlerGroup
}

// applicationId may be blank
// if blank, no applicationId validation will be performed
func NewSkill(applicationId string) *Skill {
	return &Skill{
		applicationId: applicationId,
		handlers:      HandlerGroup{},
	}
}

func (s *Skill) RegisterHandlers(handler ...RequestHandler) {
	if s.handlers == nil {
		s.handlers = HandlerGroup{}
	}
	s.handlers = append(s.handlers, handler...)
}

// ServeHTTP burasinin amaci yazilacak
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func (s *Skill) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	if os.Getenv("APP_ENV") == "production" {
		if err := validateAlexaRequest(w, r); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	requestJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if os.Getenv("GOALEXA_DUMP") != "" {
		trash := map[string]any{}
		json.Unmarshal(requestJson, &trash)
		var requestJsonPretty []byte
		if os.Getenv("GOALEXA_DUMP") == "full" {
			fmt.Println("girer")

			requestJsonPretty, _ = json.MarshalIndent(trash, "", "    ")

			fmt.Printf("MarshalIndent funnction output\n %s\n", string(requestJsonPretty))

		} else {
			fmt.Println("girer2")
			requestJsonPretty, _ = json.MarshalIndent(trash["request"], "", "    ")
		}
		Logger.Debug(fmt.Sprintf("-> -> -> From Alexa: %s", string(requestJsonPretty)))
	}

	var root alexaapi.RequestRoot
	err = json.Unmarshal(requestJson, &root)

	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		fmt.Println("ServeHTTP failed hmmm something a error", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if s.applicationId != "" && (root.Context.System.Application.ApplicationId == "" || root.Context.System.Application.ApplicationId != s.applicationId) {
		fmt.Println("Unable to verify applicationId", zap.Error(err))
		err := fmt.Errorf("Unable to verify applicationId")
		Logger.Error(
			"ServeHTTP failed",
			zap.Error(err),
			zap.String("req_skill_id", root.Context.System.Application.ApplicationId),
			zap.String("cfg_skill_id", s.applicationId),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = alexaapi.SetRequestViaLookahead(ctx, &root, requestJson)
	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		fmt.Println("ServeHTTP failed SetRequestViaLookahead", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if root.Directive.Header.Namespace != "" {
		err = alexaapi.SetEnvelopePayloadViaLookahead(ctx, &root.Directive, requestJson)
		if err != nil {
			Logger.Error("ServeHTTP failed", zap.Error(err))
			fmt.Println("ServeHTTP failed header namspace ", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	//TODO: fallback handler for when no handler takes the request

	//response, err = s.handlers.Handle(ctx, s, &root)
	//
	//if err != nil {
	//	Logger.Error("ServeHTTP failed", zap.Error(err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Println("ServeHTTP failed444", zap.Error(err))
	//	return
	//}	//if err != nil {
	//	Logger.Error("ServeHTTP failed", zap.Error(err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Println("ServeHTTP failed444", zap.Error(err))
	//	return
	//}

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
	responseJson, err := json.Marshal(response)
	_, err = io.Copy(w, bytes.NewReader(responseJson))
	//if response == nil {
	//	fmt.Println("ServeHTTP nil", zap.Error(err))
	//
	//	if os.Getenv("GOALEXA_DUMP") != "" {
	//		Logger.Debug("<- <- <- To Alexa (http 200, empty body)")
	//	}
	//
	//	w.WriteHeader(http.StatusOK)
	//	return
	//}
	//
	//responseJson, err := json.Marshal(response)
	//if err != nil {
	//	fmt.Println("ServeHTTP reee", zap.Error(err))
	//
	//	Logger.Error("ServeHTTP failed", zap.Error(err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//_, err = io.Copy(w, bytes.NewReader(responseJson))
	//
	//if err != nil {
	//	fmt.Println("ServeHTTP NewReader", zap.Error(err))
	//	Logger.Error("ServeHTTP failed", zap.Error(err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	if os.Getenv("GOALEXA_DUMP") != "" {
		responseJsonPretty, _ := json.MarshalIndent(&response, "", "    ")
		Logger.Debug(fmt.Sprintf("<- <- <- To Alexa: %s", string(responseJsonPretty)))
	}

}
