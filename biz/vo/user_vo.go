package vo

type UserStatusVO struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Department            string `json:"department"`
	TodayRecordStatus     bool   `json:"today_record_status"`
	TodayHealthCodeStatus string `json:"today_health_code_status"`
	UpdatedAt             string `json:"updated_at"`
}

type UserVO struct {
	ID                    string `json:"id"`
	Email                 string `json:"email"`
	Name                  string `json:"name"`
	Department            string `json:"department"`
	Identity              string `json:"identity"`
	Avatar                string `json:"avatar"`
	TodayRecordStatus     bool   `json:"today_record_status"`
	TodayHealthCodeStatus int    `json:"today_health_code_status"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}

type UserRecordVO struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	Address          string `json:"address"`
	TemperatureRange string `json:"temperature_range"`
	IsHealthy        bool   `json:"is_healthy"`
	HealthCodeStatus string `json:"health_code_status"`
	Remark           string `json:"remark"`
	Appendix         string `json:"appendix"`
}

type AdminVO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
