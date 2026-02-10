package database

import (
	"kasir-api-go/config"
	"kasir-api-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	var err error

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}

func Migrate() {
	err := DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.Transaction{}, &models.TransactionDetail{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}
