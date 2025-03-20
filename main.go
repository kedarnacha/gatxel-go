package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kedarnacha/gatxel-go/config"
	"github.com/kedarnacha/gatxel-go/database"
	"github.com/kedarnacha/gatxel-go/repository"
	"github.com/kedarnacha/gatxel-go/router"
	"github.com/kedarnacha/gatxel-go/service"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	fmt.Println("DB_USERNAME:", os.Getenv("DB_USERNAME"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))

	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	db := database.New(context.Background(), cfg)
	if db == nil {
		log.Fatal("Failed to initialize database connection")
	}
	log.Println("Database connected successfully")

	err = database.Migrate(cfg)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migration completed")

	r := gin.Default()

	router.SetupAppoinmentRouter(r, db)
	router.SetupNotificationRouter(r, db)
	router.SetupUserRouter(r, db)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	router.SetupAuthRouter(r, authService.(*service.AuthService))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Server is running")
	})

	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
