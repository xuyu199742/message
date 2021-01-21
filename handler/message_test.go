package handler

import (
	"message/enums"
	"message/models"
	"message/protocols"
	"context"
	"testing"
)

func TestMessage_SendText(t *testing.T) {
	type args struct {
		ctx context.Context
		req *protocols.SendTextRequest
		rsp *protocols.BaseResponse
	}
	var rsp = &protocols.BaseResponse{}
	method := make([]enums.MsgMethod, 0)
	method = append(method, "dd")
	rec := make([]string, 0)
	rec = append(rec, "321321")
	Sender := models.Sender{
		Type: 0,
		Val:  "3211321",
	}
	Receiver := models.Receiver{
		Type: 0,
		Val:  rec,
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.TODO(),
				req: &protocols.SendTextRequest{
					SendBaseRequest: protocols.SendBaseRequest{
						Sender:     Sender,
						Receiver:   Receiver,
						SendMethod: method,
					},
					Content: "12321",
				},
				rsp: rsp,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Message{}
			if err := this.SendText(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("SendText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rsp.Code != 0 {
				t.Error(rsp.Message)
			}
		})
	}
}
