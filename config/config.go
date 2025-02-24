package config

import (
	"log"
	"os"
	"gorm.io/driver/postgres" //sesuaikan dengan db yang mau kamu pakai
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL") 
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	return db
}
