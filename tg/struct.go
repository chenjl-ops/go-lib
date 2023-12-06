package tg

type tgConf struct {
	TgToken          string `json:"tg_token"`
	ChatId           string `json:"chat_id"`
	Text             string `json:"text"`
	ReplyToMessageId int    `json:"reply_to_message_id"`
}

type tgResponseData struct {
	Ok          interface{} `json:"ok"`
	Description string      `json:"description"`
	ErrorCode   int         `json:"error_code"`
	Result      struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64       `json:"id"`
			IsBot     interface{} `json:"is_bot"`
			FirstName string      `json:"first_name"`
			Username  string      `json:"username"`
		} `json:"from"`
		Chat struct {
			Id    int64  `json:"id"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}
