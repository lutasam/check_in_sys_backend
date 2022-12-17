package model

import (
	"gorm.io/gorm"
	"time"
)

type Notice struct {
	ID        uint64         `gorm:"column:id" redis:"id" json:"id"`
	UserID    uint64         `gorm:"column:user_id" redis:"user_id" json:"user_id"`
	User      User           `gorm:"foreignKey:user_id;references:id" redis:"user" json:"user"`
	Content   string         `gorm:"column:content" redis:"content" json:"content"`
	CreatedAt time.Time      `gorm:"column:created_at" redis:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" redis:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" redis:"deleted_at" json:"deleted_at"`
}

func (Notice) TableName() string {
	return "notices"
}
