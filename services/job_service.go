package services

import (
    "log"
    "scheduler-microservice/database"
    "scheduler-microservice/models"
)

func GetAllJobs() []models.Job {
    var jobs []models.Job
    if err := database.DB.Find(&jobs).Error; err != nil {
        log.Println("GetAllJobs error:", err)
    }
    return jobs
}

func GetJobByID(id string) (models.Job, bool) {
    var job models.Job
    if err := database.DB.First(&job, "id = ?", id).Error; err != nil {
        return job, false
    }
    return job, true
}

func CreateJob(job models.Job) {
    if err := database.DB.Create(&job).Error; err != nil {
        log.Println("CreateJob error:", err)
    }
}
