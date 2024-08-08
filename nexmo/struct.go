package nexmo

type Nexmo struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	AuthCode  string `json:"auth_code"`
}

type SmsSuccessResponse struct {
	MessageUUID string `json:"message_uuid"`
}

type SmsFailureResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}
