package service

import "message/models"

type WX struct {
	Mes *models.Message
}

func (W *WX) Send() (*MessageResponse, error) {
	panic("implement me")
}

func (W *WX) HandleResponse(rsp *MessageResponse, Channel *models.MessageChannel) error {
	panic("implement me")
}
