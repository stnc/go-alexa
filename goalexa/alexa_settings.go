package goalexa

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aivahealth/goalexa/alexaapi"
	"go.uber.org/zap"
)

// Helpers for accessing the Alexa Settings API
// https://developer.amazon.com/en-US/docs/alexa/smapi/alexa-settings-api-reference.html

func GetDeviceTimeZone(ctx context.Context, reqRoot *alexaapi.RequestRoot) (string, error) {
	apiEndpoint, apiAccessToken, deviceId, err := alexaApiGetCredentials(ctx, reqRoot)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/v2/devices/%s/settings/System.timeZone", apiEndpoint, deviceId), nil)
	if err != nil {
		Logger.Error("GetDeviceTimeZone failed", zap.Error(err))
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiAccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Logger.Error("GetDeviceTimeZone failed", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.Error("GetDeviceTimeZone failed", zap.Error(err))
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("expected http 200, got http %d", resp.StatusCode)
		Logger.Error("GetDeviceTimeZone failed", zap.Error(err), zap.Any("resp_body", string(bodyBytes)))
		return "", err
	}
	return string(bodyBytes), nil
}

func alexaApiGetCredentials(
	ctx context.Context,
	reqRoot *alexaapi.RequestRoot,
) (
	apiEndpoint string,
	apiAccessToken string,
	deviceId string,
	err error,
) {
	if reqRoot == nil {
		err = fmt.Errorf("reqRoot is nil")
		Logger.Error("alexaApiGetCredentials failed", zap.Error(err))
		return
	}
	apiEndpoint = reqRoot.Context.System.ApiEndpoint
	if apiEndpoint == "" {
		err = fmt.Errorf("reqRoot.Context.System.ApiEndpoint is empty")
		Logger.Error("alexaApiGetCredentials failed", zap.Error(err))
		return
	}
	apiAccessToken = reqRoot.Context.System.ApiAccessToken
	if apiAccessToken == "" {
		err = fmt.Errorf("reqRoot.Context.System.ApiAccessToken is empty")
		Logger.Error("alexaApiGetCredentials failed", zap.Error(err))
		return
	}
	deviceId = reqRoot.Context.System.Device.DeviceId
	if deviceId == "" {
		err = fmt.Errorf("reqRoot.Context.System.Device.DeviceId is empty")
		Logger.Error("alexaApiGetCredentials failed", zap.Error(err))
		return
	}
	return
}
