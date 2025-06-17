package routes

import (
    "github.com/gin-gonic/gin"
    "scheduler-microservice/controllers"
)

func SetupRoutes(r *gin.Engine) {
    jobs := r.Group("/jobs")
    {
        jobs.GET("/", controllers.GetAllJobs)
        jobs.GET("/:id", controllers.GetJobByID)
        jobs.POST("/", controllers.CreateJob)
    }
}
