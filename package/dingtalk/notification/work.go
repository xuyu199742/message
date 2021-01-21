package notification

import (
	"message/enums"
	"message/package/dingtalk"
	"github.com/xinzf/gokit/logger"

	//"message/service"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type MessageText struct {
	MessageType enums.MsgContType `json:"msgtype"`
	Text        struct {
		Content string `json:"content"`
	} `json:"text"`
}

type MessageLink struct {
	MessageType enums.MsgContType `json:"msgtype"`
	Link        struct {
		MessageUrl string `json:"messageUrl"`
		PicUrl     string `json:"picUrl"`
		Title      string `json:"title"`
		Text       string `json:"text"`
	} `json:"link"`
}

type MessageActionCard struct {
	MessageType enums.MsgContType `json:"msgtype"`
	ActionCard  struct {
		Title          string        `json:"title"`
		Markdown       string        `json:"markdown"`
		BtnOrientation string        `json:"btn_orientation"`
		BtnJsonList    []CardBtnList `json:"btn_json_list"`
	} `json:"action_card"`
}

type CardBtnList struct {
	Title     string `json:"title"`
	ActionUrl string `json:"action_url"`
}

type WorkMessageResponse struct {
	ErrCode   int    `json:"errcode"`
	TaskId    int64  `json:"task_id"`
	RequestId string `json:"request_id"`
	ErrMsg    string `json:"errmsg"`
}

type WorkReceiver struct {
	Type  enums.ReceiverType `json:"type"`
	Value string             `json:"value"`
}

type Work struct {
}

func (this *Work) send(agentId int64, msg interface{}, receiver WorkReceiver) *WorkMessageResponse {
	var (
		rsp = &WorkMessageResponse{}
	)
	req := make(map[string]interface{})
	req["agent_id"] = agentId
	req["msg"] = msg
	logger.DefaultLogger.Debug("msg", msg)
	if receiver.Type == enums.USER_RECEIVER {
		req["userid_list"] = receiver.Value
	} else {
		req["dept_id_list"] = receiver.Value
	}

	accessToken, err := dingtalk.AccessToken.GetToken()
	if err != nil {
		rsp.ErrCode = 9999
		rsp.ErrMsg = err.Error()
		return rsp
	}
	_url := fmt.Sprintf("%s/topapi/message/corpconversation/asyncsend_v2?access_token=%s", dingtalk.ACCESS_URL, accessToken)

	httpResp, _, errs := gorequest.New().Post(_url).
		Set("Content-Type", "application/json").
		SendMap(req).
		EndStruct(&rsp)

	if len(errs) > 0 {
		rsp.ErrCode = 9999
		rsp.ErrMsg = errs[0].Error()
		return rsp
	}

	if httpResp.StatusCode != 200 {
		rsp.ErrCode = 9999
		rsp.ErrMsg = fmt.Sprintf("钉钉服务器异常,httpCode: %d", httpResp.StatusCode)

		return rsp
	}

	return rsp

}

func (this Work) SendText(sender int64, content string, receiver WorkReceiver) *WorkMessageResponse {
	textObj := MessageText{}
	textObj.Text.Content = content
	textObj.MessageType = enums.TEXT
	return this.send(sender, textObj, receiver)

}

func (this Work) SendLink(sender int64, link MessageLink, receiver WorkReceiver) *WorkMessageResponse {

	return this.send(sender, link, receiver)
}

func (this Work) SendCard(sender int64, card MessageActionCard, receiver WorkReceiver) *WorkMessageResponse {
	return this.send(sender, card, receiver)
}
