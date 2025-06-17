package controllers

import (
    "net/http"
    "scheduler-microservice/models"
    "scheduler-microservice/services"
    "github.com/gin-gonic/gin"
)

func GetAllJobs(c *gin.Context) {
    c.JSON(http.StatusOK, services.GetAllJobs())
}

func GetJobByID(c *gin.Context) {
    id := c.Param("id")
    job, found := services.GetJobByID(id)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
        return
    }
    c.JSON(http.StatusOK, job)
}

func CreateJob(c *gin.Context) {
    var job models.Job
    if err := c.ShouldBindJSON(&job); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    services.CreateJob(job)
    c.JSON(http.StatusCreated, job)
}
