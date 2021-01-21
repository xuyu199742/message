package handler

import (
	"message/enums"
	"message/models"
	"message/protocols"
	"message/service"
	"context"
)

type Message struct {
}

func (this *Message) SendText(ctx context.Context, req *protocols.SendTextRequest, rsp *protocols.BaseResponse) error {
	if err := req.Error(); err != nil {
		return err
	}

	message, err := models.NewMessageExecutor().CreateMessage(req.Receiver, req.Sender, req.Msg, req.SendMethod, enums.TEXT)

	if err != nil {
		return err
	}

	service.SendMessage(message)

	return nil
}

func (this *Message) SendLink(ctx context.Context, req *protocols.SendLinkRequest, rsp *protocols.BaseResponse) error {
	if err := req.Error(); err != nil {
		return err
	}

	message, err := models.NewMessageExecutor().CreateMessage(req.Receiver, req.Sender, req.Msg, req.SendMethod, enums.LINK)

	if err != nil {
		return err
	}

	service.SendMessage(message)

	return nil
}

func (this *Message) SendCard(ctx context.Context, req *protocols.SendCardRequest, rsp *protocols.BaseResponse) error {
	if err := req.Error(); err != nil {
		return err
	}

	message, err := models.NewMessageExecutor().CreateMessage(req.Receiver, req.Sender, req.Msg, req.SendMethod, enums.CARD)

	if err != nil {
		return err
	}

	service.SendMessage(message)

	return nil
}
