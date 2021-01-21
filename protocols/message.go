package protocols

import (
	"message/enums"
	"message/models"
	"errors"
)

type SendBaseRequest struct {
	Sender     models.Sender     `json:"sender"`
	Receiver   models.Receiver   `json:"receiver"`
	SendMethod []enums.MsgMethod `json:"send_method"`
}
type SendMessageCfg struct {
	Msg models.MsgContent `json:"-"`
}

type SendTextRequest struct {
	SendBaseRequest
	SendMessageCfg
	Content string `json:"content"`
}

type SendLinkRequest struct {
	SendBaseRequest
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"pic_url"`
	MessageUrl string `json:"message_url"`
	SendMessageCfg
}

type SendCardRequest struct {
	SendBaseRequest
	Title   string `json:"title"`
	Content string `json:"content"`
	BtnList []struct {
		Title  string `json:"title"`
		ButUrl string `json:"but_url"`
	} `json:"btn_list"`
	SendMessageCfg
}

func (this *SendBaseRequest) Error() error {
	if len(this.SendMethod) < 0 {
		return errors.New("发送方式不能为空")
	}

	if this.Sender.Val == "" {
		return errors.New("发送人不能为空")
	}
	if len(this.Receiver.Val) < 0 {
		return errors.New("接收人不能为空")
	}

	if !this.Receiver.Type.Is() {
		return errors.New("非法receiver类型")
	}
	if !this.Sender.Type.Is() {
		return errors.New("非法sender类型")
	}

	for _, v := range this.SendMethod {
		if !v.Is() {
			return errors.New("非法send method")
		}
	}

	return nil
}

func (this *SendTextRequest) Error() error {
	if this.Content == "" {
		return errors.New("text消息内容不能为空")
	}
	this.Msg = map[string]interface{}{
		"content": this.Content,
	}

	return nil
}

func (this *SendLinkRequest) Error() error {
	if this.Text == "" || this.Title == "" || this.PicUrl == "" || this.MessageUrl == "" {
		return errors.New("link消息缺少参数")
	}
	this.Msg = map[string]interface{}{
		"text":        this.Text,
		"pic_url":     this.PicUrl,
		"title":       this.Title,
		"message_url": this.MessageUrl,
	}
	return nil
}

func (this *SendCardRequest) Error() error {
	if this.Title == "" || this.Content == "" || len(this.BtnList) < 0 {
		return errors.New("card消息缺少参数")
	}

	this.Msg = map[string]interface{}{
		"title":    this.Title,
		"content":  this.Content,
		"btn_list": this.BtnList,
	}

	return nil
}
