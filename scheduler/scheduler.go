package scheduler

import (
	"fmt"
	"scheduler-microservice/database"
	"scheduler-microservice/models"
	"time"

	"github.com/robfig/cron/v3"
)

// StartCronScheduler initializes the scheduler and registers jobs using cron expressions.
func StartCronScheduler() {
	var jobs []models.Job
	if err := database.DB.Debug().Find(&jobs).Error; err != nil {
		fmt.Println("Error fetching jobs from DB:", err)
		return
	}

	c := cron.New()

	for _, job := range jobs {
		jobCopy := job // avoid closure capture issue

		_, err := c.AddFunc(job.CronExpression, func() {
			fmt.Printf("Running job [%s] at %s\n", jobCopy.Name, time.Now().Format(time.RFC3339))
			dummyNumberCrunching(jobCopy)

			// update last run time
			database.DB.Model(&jobCopy).Update("last_run_at", time.Now())
		})

		if err != nil {
			fmt.Printf("⚠️ Failed to schedule job [%s]: %v\n", jobCopy.Name, err)
		}
	}

	c.Start()

	fmt.Println("✅ Cron scheduler started and jobs registered.")
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
