package nexmo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/chenjl-ops/go-lib/requests"
)

func New(apiKey string, apiSecret string) *Nexmo {
	authCode := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", apiKey, apiSecret)))

	return &Nexmo{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		AuthCode:  authCode,
	}
}

func (s *SmsSuccessResponse) validate() error {
	if s.MessageUUID == "" {
		return errors.New("message uuid is empty")
	}
	return nil
}

func (n *Nexmo) GetAuthCode() string {
	return n.AuthCode
}

func (n *Nexmo) GetApiKey() string {
	return n.ApiKey
}

func (n *Nexmo) GetApiSecret() string {
	return n.ApiSecret
}

func (n *Nexmo) GetAuthHeaders() map[string]string {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", n.AuthCode),
		"Content-Type":  "application/json",
	}
	return headers
}

func (n *Nexmo) SendSms(from string, to string, message string) (map[string]interface{}, error) {
	url := "https://api.nexmo.com/v1/messages"

	requestData := map[string]string{
		"message_type": "text",
		"text":         message,
		"channel":      "sms",
		"from":         from,
		"to":           to,
	}

	var successData SmsSuccessResponse
	var failureData SmsFailureResponse
	err := requests.RequestWithError(url, "POST", n.GetAuthHeaders(), requestData, &successData, &failureData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var data map[string]interface{}

	if err := successData.validate(); err != nil {
		data["error"] = err.Error()
		data["success"] = false
		data["data"] = failureData
	} else {
		data["success"] = true
		data["data"] = successData
	}
	return data, nil
}
