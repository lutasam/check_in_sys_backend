package vo

type RecordVO struct {
	ID               string `json:"id"`
	Address          string `json:"address"`
	TemperatureRange string `json:"temperature_range"`
	IsHealthy        bool   `json:"is_healthy"`
	HealthCodeStatus string `json:"health_code_status"`
	CreatedAt        string `json:"created_at"`
}
