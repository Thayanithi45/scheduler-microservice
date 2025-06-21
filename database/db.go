package database

import (
	"log"
	"os"
	"time"

	"scheduler-microservice/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DATABASE_URL")
	var db *gorm.DB
	var err error

	// Retry DB connection up to 5 times
	for i := 1; i <= 5; i++ {
		log.Printf("⏳ Connecting to DB (attempt %d/5)...", i)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Connected to DB successfully")
			break
		}

		log.Printf("❌ Failed to connect: %v", err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to database after 5 attempts: %v", err)
	}

	// Run migrations
	if err := db.AutoMigrate(&models.Job{}); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	DB = db
	log.Println("Database connected and migrated successfully")
}
