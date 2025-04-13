package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("❌ Error occured while reading .env file")
	}

	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatal("❌ DB_URL not found")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Could not connect to the database")
	}

	fmt.Println("🚀 Connected to Database successfully")

	return db
}
