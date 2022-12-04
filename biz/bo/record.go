package bo

import "github.com/lutasam/check_in_sys/biz/vo"

type UploadUserRecordRequest struct {
	Address          string `json:"address" binding:"required"`
	TemperatureRange int    `json:"temperature_range" binding:"required"`
	IsHealthy        bool   `json:"is_healthy" binding:"required"`
	HealthCodeStatus int    `json:"health_code_status" binding:"required"`
	Remark           string `json:"remark" binding:"-"`
	Appendix         string `json:"appendix" binding:"-"`
}

type UploadUserRecordResponse struct{}

type FindUserAllRecordsRequest struct {
	CurrentPage int `json:"current_page" binding:"required"`
	PageSize    int `json:"page_size" binding:"required"`
}

type FindUserAllRecordsResponse struct {
	Total   int            `json:"total"`
	Records []*vo.RecordVO `json:"records"`
}

type FindUserTodayRecordRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type FindUserTodayRecordResponse struct {
	UserRecord *vo.UserRecordVO `json:"user_record"`
}
