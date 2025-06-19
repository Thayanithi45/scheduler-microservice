package models

import "time"

type Job struct {
	ID             int       `json:"id" gorm:"primaryKey;auto_increment;type:serial"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	CronExpression string    `json:"cron_expression"`
	LastRunAt      time.Time `json:"last_run_at"`
	NextRunAt      time.Time `json:"next_run_at"`
	Status         string    `json:"status"`
}
