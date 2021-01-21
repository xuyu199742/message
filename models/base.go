package models

import (
	"gorm.io/gorm"
)

type Model struct {
	ID      int   `gorm:"primarykey" json:"id"`
	Updated int64 `gorm:"autoUpdateTime:milli" json:"updated"`
	Created int64 `gorm:"autoCreateTime:milli" json:"created"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
