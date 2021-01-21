package models

import (
	"message/enums"
	"github.com/jinzhu/gorm"
	"github.com/xinzf/gokit/storage"
	"github.com/xinzf/gokit/utils"
)

type MessageChannel struct {
	Model
	MessageID  int64           `gorm:"column:message_id;type:int" json:"message_id"`
	Status     enums.MsgStatus `gorm:"column:status;type:tinyint" json:"status"`
	Response   utils.GormMap   `gorm:"column:response;type:json" json:"response"`
	Request    MsgContent      `gorm:"column:request;type:json" json:"request"`
	SendMethod enums.MsgMethod `gorm:"column:send_method;type:varchar(50)" json:"send_method"`
}

type MessageChannelExecutor struct {
	db *gorm.DB
}

func NewMessageChannelExecutor(db ...*gorm.DB) *MessageChannelExecutor {
	if len(db) > 0 {
		return &MessageChannelExecutor{db: db[0]}
	}
	return &MessageChannelExecutor{db: storage.DB.Use()}
}

func (this *MessageChannelExecutor) Save(channel *MessageChannel, columns ...[]string) error {
	if len(columns) > 0 {
		return this.db.Select(columns).Save(channel).Error
	} else {
		return this.db.Save(channel).Error
	}
}
