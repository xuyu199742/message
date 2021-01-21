package handler

import (
	"message/enums"
	"message/protocols"
	"context"
)

type Enums struct {
}

func (*Enums) List(c context.Context, request *protocols.EnumsTypeRequest, response *protocols.BaseResponse) error {

	res := make(map[string]interface{})
	{
		res["msg_method"] = enums.MsgMethodToOptions()
		res["sender"] = enums.SenderTypeToOptions()
		res["receiver"] = enums.ReceiverTypeToOptions()
	}

	if request.Type != "" {
		ret, found := res[request.Type]
		if found {
			response.Result = ret
			return nil
		}
	}

	response.Result = res
	return nil
}
