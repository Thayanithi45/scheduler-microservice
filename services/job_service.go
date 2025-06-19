package services

import (
	"errors"
	"scheduler-microservice/database"
	"scheduler-microservice/models"
	"scheduler-microservice/scheduler"
	"time"

	"github.com/robfig/cron/v3"
)

func GetAllJobs() ([]models.Job, error) {
	var jobs []models.Job
	err := database.DB.Find(&jobs).Error
	return jobs, err
}

func GetJobByID(id string) (models.Job, error) {
	var job models.Job
	err := database.DB.First(&job, "id = ?", id).Error
	if err != nil {
		return job, errors.New("job not found")
	}
	return job, nil
}

func CreateJob(job models.Job) (models.Job, error) {
	// Generate ID
	job.Status = "scheduled"
	job.LastRunAt = time.Time{}

	// Validate cron expression
	sched, err := cron.ParseStandard(job.CronExpression)
	if err != nil {
		return job, errors.New("invalid cron expression")
	}

	// Compute next run
	nextRun := sched.Next(time.Now())
	job.NextRunAt = nextRun

	// Save job
	if err := database.DB.Create(&job).Error; err != nil {
		return job, err
	}

	// Register job in scheduler
	scheduler.RegisterJob(job)

	return job, nil
}
