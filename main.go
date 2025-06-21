package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
    "scheduler-microservice/database"
    "scheduler-microservice/routes"
    "scheduler-microservice/scheduler"
        _ "scheduler-microservice/docs" // ✅ ADD THIS LINE


    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// @title Scheduler Microservice API
// @version 1.0
// @description API for managing and scheduling jobs
// @host localhost:8080
// @BasePath /
func main() {
    // Load .env
    // Only load .env if not running in Docker
    if os.Getenv("RUN_ENV") != "docker" {
        if err := godotenv.Load(".env"); err != nil {
            log.Println("No .env file found")
        }
    }

    // Init DB and Scheduler
    database.Init()
	scheduler.StartCronScheduler()

    // Start Gin
    r := gin.Default()
    
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    routes.SetupRoutes(r)
    r.Run(os.Getenv("PORT"))
}
