package database

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/kedarnacha/gatxel-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(ctx context.Context, cfg config.Config) *gorm.DB {
	pass := url.QueryEscape(cfg.DatabasePassword)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DatabaseHost,
		cfg.DatabaseUsername,
		pass,
		cfg.DatabaseName,
		cfg.DatabasePort,
	)

	log.Printf("Connecting to DB with: host=%s user=%s dbname=%s port=%s",
		cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabaseName, cfg.DatabasePort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Printf(" Connected to the database")

	return db
}
