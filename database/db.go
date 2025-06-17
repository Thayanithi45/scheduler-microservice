package database

import (
    "log"
    "os"
    "scheduler-microservice/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("DB connection failed: %v", err)
    }

    if err := db.AutoMigrate(&models.Job{}); err != nil {
        log.Fatalf("AutoMigrate error: %v", err)
    }

    DB = db
    log.Println("Database connected and migrated")
}
