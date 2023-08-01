package goalexa

import (
	"encoding/json"

	"go.uber.org/zap"
)

//
//
// Helpers for session attributes

//
//
// Import/Export
// Using JSON as an intermediary, we can import/export well-typed data

func ExportAttributes[T any](attrOut *T) (map[string]any, error) {
	marshaled, err := json.Marshal(&attrOut)
	if err != nil {
		Logger.Error("ExportAttributes failed", zap.Error(err))
		return nil, err
	}
	result := map[string]any{}
	err = json.Unmarshal(marshaled, &result)
	if err != nil {
		Logger.Error("ExportAttributes failed", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func ImportAttributes[T any](attrIn map[string]any) (*T, error) {
	marshaled, err := json.Marshal(&attrIn)
	if err != nil {
		Logger.Error("ImportAttributes failed", zap.Error(err))
		return nil, err
	}
	var result T
	err = json.Unmarshal(marshaled, &result)
	if err != nil {
		Logger.Error("ImportAttributes failed", zap.Error(err))
		return nil, err
	}
	return &result, nil
}
