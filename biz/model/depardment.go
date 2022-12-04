package model

import (
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID        uint64         `gorm:"column:id"`
	Name      string         `gorm:"name"`
	Admin     User           `gorm:"foreignKey:id;references:admin_id"`
	Users     []User         `gorm:"foreignKey:department_id"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
