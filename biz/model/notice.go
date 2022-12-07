package model

import (
	"gorm.io/gorm"
	"time"
)

type Notice struct {
	ID        uint64         `gorm:"column:id"`
	UserID    uint64         `gorm:"column:user_id"`
	User      User           `gorm:"foreignKey:user_id;references:id""`
	Content   string         `gorm:"column:content"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Notice) TableName() string {
	return "notices"
}
