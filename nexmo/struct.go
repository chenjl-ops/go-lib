package nexmo

type Nexmo struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	AuthCode  string `json:"auth_code"`
}

type SmsMessageSuccessResponse struct {
	MessageUUID string `json:"message_uuid"`
}

type SmsMessageFailureResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type SmsBaseData struct {
	To               string `json:"to"`
	MessageId        string `json:"message-id"`
	Status           string `json:"status"`
	RemainingBalance string `json:"remaining-balance"`
	MessagePrice     string `json:"message-price"`
	Network          string `json:"network"`
	ClientRef        string `json:"client-ref"`
	AccountRef       string `json:"account-ref"`
	ErrorText        string `json:"error-text"`
}

type SmsBaseResponse struct {
	MessageCount string        `json:"message-count"`
	Messages     []SmsBaseData `json:"messages"`
}
