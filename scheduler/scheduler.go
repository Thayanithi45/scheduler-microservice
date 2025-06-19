package scheduler

import (
	"fmt"
	"scheduler-microservice/database"
	"scheduler-microservice/models"
	"time"

	"github.com/robfig/cron/v3"
)

var CronScheduler *cron.Cron

// StartCronScheduler initializes the scheduler and registers jobs using cron expressions.
func StartCronScheduler() {
	CronScheduler = cron.New()
	loadJobsFromDB()
	CronScheduler.Start()
}

// loadJobsFromDB fetches jobs from the DB and registers them with the scheduler.
func loadJobsFromDB() {
	var jobs []models.Job
	if err := database.DB.Find(&jobs).Error; err != nil {
		fmt.Println("Error fetching jobs from DB:", err)
		return
	}

	for _, job := range jobs {
		RegisterJob(job)
	}
}

// registerJob registers a single job with the scheduler.
func RegisterJob(job models.Job) {
	jobCopy := job // avoid closure capture issue

	_, err := CronScheduler.AddFunc(job.CronExpression, func() {
		fmt.Printf("⏰ Running job [%s] at %s\n", jobCopy.Name, time.Now().Format(time.RFC3339))

		// Dummy execution logic
		dummyNumberCrunching(jobCopy)

		now := time.Now()
		database.DB.Model(&jobCopy).Updates(map[string]interface{}{
			"last_run_at": now.Format(time.RFC3339),
		})
	})

	if err != nil {
		fmt.Printf("Failed to register job [%s]: %v\n", job.Name, err)
	}
}
func dummyNumberCrunching(job models.Job) {
	fmt.Printf("🔢 Starting number crunching for job: %s\n", job.Name)
	var sum int64
	for i := int64(1); i <= 1000000; i++ {
		sum += i
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("✅ Finished job [%s] → Result: %d\n", job.Name, sum)
}
