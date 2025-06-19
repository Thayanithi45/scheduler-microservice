package controllers

import (
	"net/http"
	"scheduler-microservice/models"
	"scheduler-microservice/services"

	"github.com/gin-gonic/gin"
)
// GetAllJobs godoc
// @Summary      List all jobs
// @Description  Get all scheduled jobs from the database
// @Tags         jobs
// @Produce      json
// @Success      200  {array}  models.Job
// @Router       /jobs [get]
// GetJobs lists all scheduled jobs
func GetJobs(c *gin.Context) {
	jobs, err := services.GetAllJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}


// GetJobByID godoc
// @Summary      Get a job by ID
// @Description  Get job details by its ID
// @Tags         jobs
// @Produce      json
// @Param        id   path      string  true  "Job ID"
// @Success      200  {object}  models.Job
// @Failure      404  {object}  map[string]string
// @Router       /jobs/{id} [get]
// GetJobByID retrieves a job by ID
func GetJobByID(c *gin.Context) {
	id := c.Param("id")
	job, err := services.GetJobByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}


// CreateJob godoc
// @Summary      Create a new job
// @Description  Add a new scheduled job to the database
// @Tags         jobs
// @Accept       json
// @Produce      json
// @Param        job  body      models.Job  true  "Job to create"
// @Success      201  {object}  models.Job
// @Failure      400  {object}  map[string]string
// @Router       /jobs [post]
// CreateJob handles job creation
func CreateJob(c *gin.Context) {
	var input models.Job
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdJob, err := services.CreateJob(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdJob)
}
