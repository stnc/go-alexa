package goalexa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aivahealth/goalexa/alexaapi"
	"go.uber.org/zap"
	"io"
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

	requestJson, err := io.ReadAll(r.Body)

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
			requestJsonPretty, _ = json.MarshalIndent(trash, "", "    ")
		} else {
			requestJsonPretty, _ = json.MarshalIndent(trash["request"], "", "    ")
		}
		Logger.Debug(fmt.Sprintf("-> -> -> From Alexa: %s", string(requestJsonPretty)))
	}

	var root alexaapi.RequestRoot
	err = json.Unmarshal(requestJson, &root)

	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if s.applicationId != "" && (root.Context.System.Application.ApplicationId == "" || root.Context.System.Application.ApplicationId != s.applicationId) {
		err := fmt.Errorf("Unable to verify applicationId")
		Logger.Error(
			"ServeHTTP failed",
			zap.Error(err),
			zap.String("req_skill_id", root.Context.System.Application.ApplicationId),
			zap.String("cfg_skill_id", s.applicationId),
		)
		//fmt.Println("app id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = alexaapi.SetRequestViaLookahead(ctx, &root, requestJson)
	if err != nil {
		//fmt.Println("SetRequestViaLookahead")
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if root.Directive.Header.Namespace != "" {
		err = alexaapi.SetEnvelopePayloadViaLookahead(ctx, &root.Directive, requestJson)
		if err != nil {
			//fmt.Println("SetEnvelopePayloadViaLookahead")
			Logger.Error("ServeHTTP failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// TODO: fallback handler for when no handler takes the request
	response, err := s.handlers.Handle(ctx, s, &root)
	if err != nil {
		//fmt.Println("handlers.Handl")
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if response == nil {
		if os.Getenv("GOALEXA_DUMP") != "" {
			Logger.Debug("<- <- <- To Alexa (http 200, empty body)")
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		//fmt.Println("json.Marshal(response)")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: why?
	_, err = io.Copy(w, bytes.NewReader(responseJson))
	if err != nil {
		//fmt.Println("bytes.NewReade")
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if os.Getenv("GOALEXA_DUMP") != "" {
		responseJsonPretty, _ := json.MarshalIndent(&response, "", "    ")
		Logger.Debug(fmt.Sprintf("<- <- <- To Alexa: %s", string(responseJsonPretty)))
	}
}
