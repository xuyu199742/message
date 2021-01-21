package service

import (
	"message/enums"
	"message/models"
	"github.com/xinzf/gokit/utils"
)

type Sender interface {
	Send() *MessageResponse
	HandleResponse(rsp *MessageResponse, Channel *models.MessageChannel)
}

type Transformer interface {
	Text(body []byte) (Text, error)
	Link(body []byte) (Link, error)
	Card(body []byte) (Card, error)
}

type Text struct {
	Content string `json:"content"`
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"pic_url"`
	MessageUrl string `json:"message_url"`
}
type Card struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	BtnList []struct {
		Title  string `json:"title"`
		ButUrl string `json:"but_url"`
	} `json:"btn_list"`
}

type MessageResponse struct {
	ErrCode int           `json:"errcode"`
	Res     utils.GormMap `json:"res"`
}

func SendMessage(message *models.Message) {
	var rsp = &MessageResponse{}
	for _, channel := range message.Channels {
		go func(msChan *models.MessageChannel) {
			ms := getMsObj(msChan.SendMethod, message)
			rsp = ms.Send()
			ms.HandleResponse(rsp, msChan)
		}(channel)
	}
}

func getMsObj(typ enums.MsgMethod, mes *models.Message) Sender {
	switch typ {
	case enums.DD:
		var dd DD
		dd.Mes = mes
		return &dd
	default:
		return nil
	}
}
