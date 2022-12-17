package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                    uint64         `gorm:"column:id" redis:"id" json:"id"`
	Email                 string         `gorm:"column:email" redis:"email" json:"email"`
	Password              string         `gorm:"column:password" redis:"password" json:"password"`
	Name                  string         `gorm:"column:name" redis:"name" json:"name"`
	DepartmentID          uint64         `gorm:"column:department_id" redis:"department_id" json:"department_id"`
	Avatar                string         `gorm:"column:avatar" redis:"avatar" json:"avatar"`
	TodayRecordStatus     bool           `gorm:"column:today_record_status" redis:"today_record_status" json:"today_record_status"`
	TodayHealthCodeStatus int            `gorm:"column:today_health_code_status" redis:"today_health_code_status" json:"today_health_code_status"`
	Identity              int            `gorm:"column:identity" redis:"identity" json:"identity"`
	CreatedAt             time.Time      `gorm:"column:created_at" redis:"created_at" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"column:updated_at" redis:"updated_at" json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"column:deleted_at" redis:"deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
