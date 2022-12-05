package model

import (
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID        uint64         `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	AdminID   uint64         `gorm:"column:admin_id"`
	Admin     User           `gorm:"foreignKey:admin_id;references:id"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Department) TableName() string {
	return "department"
}
