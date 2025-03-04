package main

import (
	"context"
	"log"
	"net/http"

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
	cfg, err := env.ParseAs[config.config]()
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	// Initialize database connection
	db, err := database.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	err = database.Migrate(cfg)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Setup routers
	router.SetupAppoinmentRouter(r, db)
	router.SetupNotificationRouter(r, db)
	router.SetupUserRouter(r, db)

	// Initialize authentication components
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
