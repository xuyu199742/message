package models

import (
	"message/enums"
	"database/sql/driver"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"github.com/xinzf/gokit/storage"
	"github.com/xinzf/gokit/utils"
)

type Message struct {
	Model
	UUID        string            `gorm:"column:uuid;type:varchar(255)" json:"uuid"`
	Receiver    Receiver          `gorm:"column:receiver;type:json" json:"receiver"`
	MsgContType enums.MsgContType `gorm:"column:msg_cont_type;type:varchar(50)" json:"msg_cont_type"`
	MsgType     enums.MsgType     `gorm:"column:msg_type;type:varchar(50)" json:"msg_type"`
	Sender      Sender            `gorm:"column:sender;type:json" json:"sender"`
	MsgContent  MsgContent        `gorm:"column:content;type:json" json:"msg_content"`
	Channels    []*MessageChannel `json:"channel"`
}

type Receiver struct {
	Type enums.ReceiverType ` json:"type"`
	Val  []string           `json:"value"`
}

type MsgContent utils.GormMap

type Sender struct {
	Type enums.SenderType `json:"type"`
	Val  string           `json:"value"`
}

func (static Sender) Value() (driver.Value, error) {
	return jsoniter.MarshalToString(static)
}

func (this *Sender) Scan(v interface{}) error {
	var strs Sender
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}
	*this = strs
	return nil
}

func (static MsgContent) Value() (driver.Value, error) {
	return jsoniter.MarshalToString(static)
}

func (this *MsgContent) Scan(v interface{}) error {
	var strs MsgContent
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}
	*this = strs
	return nil
}

func (static Receiver) Value() (driver.Value, error) {
	return jsoniter.MarshalToString(static)
}

func (this *Receiver) Scan(v interface{}) error {
	var strs Receiver
	if err := jsoniter.Unmarshal(v.([]byte), &strs); err != nil {
		return err
	}
	*this = strs
	return nil
}

type MessageExecutor struct {
	db *gorm.DB
}

func NewMessageExecutor(db ...*gorm.DB) *MessageExecutor {
	if len(db) > 0 {
		return &MessageExecutor{db: db[0]}
	}
	return &MessageExecutor{db: storage.DB.Use()}
}

func (this *MsgContent) Marshal() ([]byte, error) {
	marshal, err := jsoniter.Marshal(this)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (this *MessageExecutor) CreateMessage(rec Receiver, sen Sender, msg MsgContent, method []enums.MsgMethod, msgContType enums.MsgContType) (*Message, error) {
	ms := &Message{
		UUID:        utils.UUID(),
		Receiver:    rec,
		Sender:      sen,
		MsgContent:  msg,
		MsgContType: msgContType,
		MsgType:     enums.WORK,
	}
	channel := MessageChannel{}
	for _, v := range method {
		channel.SendMethod = v
		channel.Request = msg
		channel.Status = enums.Sending
		ms.Channels = append(ms.Channels, &channel)
	}
	err := this.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&ms).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return ms, nil
}
