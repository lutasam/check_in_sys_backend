package model

import (
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID        uint64         `gorm:"column:id" redis:"id" json:"id"`
	Name      string         `gorm:"column:name" redis:"name" json:"name"`
	AdminID   uint64         `gorm:"column:admin_id" redis:"admin_id" json:"admin_id"`
	Admin     User           `gorm:"foreignKey:admin_id;references:id" redis:"admin" json:"admin"`
	CreatedAt time.Time      `gorm:"column:created_at" redis:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" redis:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" redis:"deleted_at" json:"deleted_at"`
}

func (Department) TableName() string {
	return "department"
}
