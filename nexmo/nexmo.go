package nexmo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/chenjl-ops/go-lib/requests"
	"slices"
	"strconv"
)

func New(apiKey string, apiSecret string) *Nexmo {
	authCode := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", apiKey, apiSecret)))

	return &Nexmo{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		AuthCode:  authCode,
	}
}

func (s *SmsMessageSuccessResponse) validate() error {
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

// SendSmsMessage 向某个渠道发送sms
func (n *Nexmo) SendSmsMessage(from string, to string, message string) (map[string]interface{}, error) {
	url := "https://api.nexmo.com/v1/messages"

	requestData := map[string]string{
		"message_type": "text",
		"text":         message,
		"channel":      "sms",
		"from":         from,
		"to":           to,
	}

	var successData SmsMessageSuccessResponse
	var failureData SmsMessageFailureResponse
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

// SendSMS 手机短信
func (n *Nexmo) SendSMS(from string, to string, message string, messageType string) (map[string]interface{}, error) {
	MessageTypes := []string{"json", "xml"}
	if slices.Contains(MessageTypes, messageType) {
		return nil, errors.New("message type is invalid: " + messageType)
	}

	url := fmt.Sprintf("https://api.nexmo.com/v1/%s", messageType)
	requestData := map[string]string{
		"from":       from,
		"to":         to,
		"text":       message,
		"type":       "unicode",
		"api_key":    n.ApiKey,
		"api_secret": n.ApiSecret,
	}

	var smsData SmsBaseResponse
	err := requests.Request(url, "POST", n.GetAuthHeaders(), requestData, &smsData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var data map[string]interface{}
	number, err := strconv.Atoi(smsData.MessageCount)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		if (number >= 1) && (smsData.Messages[0].Status == "0") {
			data["success"] = true
		} else {
			data["success"] = false
			data["error"] = smsData.Messages[0].ErrorText
		}
		data["data"] = smsData
	}
	return data, nil
}
