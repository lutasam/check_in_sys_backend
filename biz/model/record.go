package model

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	ID               uint64         `gorm:"column:id" redis:"id" json:"id"`
	UserID           uint64         `gorm:"column:user_id" redis:"user_id" json:"user_id"`
	Address          string         `gorm:"column:address" redis:"address" json:"address"`
	TemperatureRange int            `gorm:"column:temperature_range" redis:"temperature_range" json:"temperature_range"`
	IsHealthy        bool           `gorm:"column:is_healthy" redis:"is_healthy" json:"is_healthy"`
	HealthCodeStatus int            `gorm:"column:health_code_status" redis:"health_code_status" json:"health_code_status"`
	Remark           string         `gorm:"column:remark" redis:"remark" json:"remark"`
	Appendix         string         `gorm:"column:appendix" redis:"appendix" json:"appendix"`
	CreatedAt        time.Time      `gorm:"column:created_at" redis:"created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" redis:"updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" redis:"deleted_at" json:"deleted_at"`
}

func (Record) TableName() string {
	return "records"
}
