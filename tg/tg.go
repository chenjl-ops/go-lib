package tg

import (
	"fmt"
	"github.com/chenjl-ops/go-lib/requests"
)

func NewTgConf(token string, chatId string, text string, replyToMessageId int) (conf *tgConf, err error) {
	conf = &tgConf{
		TgToken:          token,
		ChatId:           chatId,
		Text:             text,
		ReplyToMessageId: replyToMessageId,
	}
	return conf, nil
}

func (tgData *tgConf) SendMessage() (error, *tgResponseData) {
	tgUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", tgData.TgToken, tgData.ChatId, tgData.Text)
	headers := map[string]string{"Content-type": "application/json"}
	data := new(tgResponseData)

	err := requests.Request(tgUrl, "GET", headers, nil, data)
	if nil != err {
		return err, nil
	} else {
		return nil, data
	}
}
