package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "scheduler-microservice/database"
    "scheduler-microservice/routes"
    "scheduler-microservice/scheduler"
    "os"
)

// @title Scheduler Microservice API
// @version 1.0
// @description API for managing and scheduling jobs
// @host localhost:8080
// @BasePath /
func main() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    // Init DB and Scheduler
    database.Init()
	scheduler.StartCronScheduler()

    // Start Gin
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(os.Getenv("PORT"))
}
