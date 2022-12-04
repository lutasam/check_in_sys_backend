package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                    uint64         `gorm:"column:id"`
	Email                 string         `gorm:"column:email"`
	Password              string         `gorm:"column:password"`
	Name                  string         `gorm:"column:name"`
	DepartmentID          uint64         `gorm:"column:department_id"`
	Avatar                string         `gorm:"column:avatar"`
	TodayRecordStatus     bool           `gorm:"column:today_record_status"`
	TodayHealthCodeStatus int            `gorm:"column:today_health_code_status"`
	Identity              int            `gorm:"column:identity"`
	Records               []Record       `gorm:"foreignKey:user_id"`
	CreatedAt             time.Time      `gorm:"column:created_at"`
	UpdatedAt             time.Time      `gorm:"column:updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
