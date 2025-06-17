package models

type Job struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	CronExpression string `json:"cron_expression"`
	LastRunAt      string `json:"last_run_at"`
	NextRunAt      string `json:"next_run_at"`
	Status         string `json:"status"`
}
