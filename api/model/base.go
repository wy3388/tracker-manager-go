package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	UUID       string `json:"uuid"`
	CreateTime string `json:"create_time" gorm:"column:CREATE_TIME"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.UUID = uuid.New().String()
	b.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
