package model

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	ID               uint64         `gorm:"column:id"`
	UserID           uint64         `gorm:"column:user_id"`
	Address          string         `gorm:"column:address"`
	TemperatureRange int            `gorm:"column:temperature_range"`
	IsHealthy        bool           `gorm:"column:is_healthy"`
	HealthCodeStatus int            `gorm:"column:health_code_status"`
	Remark           string         `gorm:"column:remark"`
	Appendix         string         `gorm:"column:appendix"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Record) TableName() string {
	return "records"
}
