package routes

import (
	"scheduler-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	
	jobs := r.Group("/jobs")
	{
		jobs.GET("/", controllers.GetJobs)
		jobs.GET("detail/:id", controllers.GetJobByID)
		jobs.POST("create/", controllers.CreateJob)
	}
}
