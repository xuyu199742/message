package service

import (
	"message/enums"
	"message/models"
	"message/package/dingtalk/notification"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"strings"
)

type DD struct {
	Mes *models.Message
}

//todo 能否拆分send消息类型
//todo err能否再优化
func (D *DD) Send() *MessageResponse {
	var (
		w         notification.Work
		seder     int64
		err       error
		rsp       = &MessageResponse{}
		res       = &notification.WorkMessageResponse{}
		wReceiver notification.WorkReceiver
	)
	map_ := make(map[string]interface{})
	if seder, err = cast.ToInt64E(D.Mes.Sender.Val); err != nil {
		rsp.ErrCode = 9999
		map_["systemErr"] = err.Error()
		rsp.Res = map_
		return rsp
	}
	wReceiver.Type = D.Mes.Receiver.Type
	wReceiver.Value = strings.Join(D.Mes.Receiver.Val, ",")
	switch D.Mes.MsgContType {
	case enums.LINK:
		body, err := D.Mes.MsgContent.Marshal()
		ms := notification.MessageLink{}
		if err != nil {
			rsp.ErrCode = 9999
			map_["systemErr"] = err.Error()
			rsp.Res = map_
			return rsp
		}
		msLink, err := D.Link(body)
		if err != nil {
			rsp.ErrCode = 9999
			map_["systemErr"] = err.Error()
			rsp.Res = map_
			return rsp
		}
		ms.MessageType = D.Mes.MsgContType
		ms.Link.MessageUrl = msLink.MessageUrl
		ms.Link.Title = msLink.Title
		ms.Link.Text = msLink.Text
		ms.Link.PicUrl = msLink.PicUrl

		res = w.SendLink(seder, ms, wReceiver)
	case enums.TEXT:
		body, err := D.Mes.MsgContent.Marshal()
		if err != nil {
			rsp.ErrCode = 9999
			map_["systemErr"] = err.Error()
			rsp.Res = map_
			return rsp
		}
		text, err := D.Text(body)
		if err != nil {
			rsp.ErrCode = 9999
			map_["systemErr"] = err.Error()
			rsp.Res = map_
			return rsp
		}
		res = w.SendText(seder, text.Content, wReceiver)
	case enums.CARD:
		body, err := D.Mes.MsgContent.Marshal()
		ms := notification.MessageActionCard{}
		msBtn := notification.CardBtnList{}
		if err != nil {
			rsp.ErrCode = 9999
			map_["systemErr"] = err.Error()
			rsp.Res = map_
			return rsp
		}
		card, err := D.Card(body)
		if err != nil {
			rsp.ErrCode = 9999
			map_["systemErr"] = err.Error()
			rsp.Res = map_
			return rsp
		}
		ms.MessageType = "action_card"
		ms.ActionCard.Title = card.Title
		ms.ActionCard.Markdown = card.Content
		ms.ActionCard.BtnOrientation = "0"
		for _, v := range card.BtnList {
			msBtn.Title = v.Title
			msBtn.ActionUrl = v.ButUrl
			ms.ActionCard.BtnJsonList = append(ms.ActionCard.BtnJsonList, msBtn)
		}
		res = w.SendCard(seder, ms, wReceiver)
	}

	rsp.ErrCode = res.ErrCode
	str, _ := jsoniter.Marshal(res)
	if rsp.Res, err = cast.ToStringMapE(string(str)); err != nil {
		rsp.ErrCode = 9999
		map_["systemErr"] = err.Error()
		rsp.Res = map_
		return rsp
	}
	return rsp
}

func (D *DD) HandleResponse(rsp *MessageResponse, channel *models.MessageChannel) {
	if rsp.ErrCode == 0 {
		channel.Status = enums.DONE
	} else {
		channel.Status = enums.FAIL
	}
	channel.Response = rsp.Res

	_ = models.NewMessageChannelExecutor().Save(channel)
}

func (D *DD) Text(body []byte) (Text, error) {
	var text = Text{}
	if err := jsoniter.Unmarshal(body, &text); err != nil {
		return text, err
	}
	return text, nil
}

func (D *DD) Link(body []byte) (Link, error) {
	var link = Link{}
	if err := jsoniter.Unmarshal(body, &link); err != nil {
		return link, err
	}
	return link, nil
}

func (D *DD) Card(body []byte) (Card, error) {
	var card = Card{}
	if err := jsoniter.Unmarshal(body, &card); err != nil {
		return card, err
	}
	return card, nil
}
