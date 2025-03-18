package main

import (
	"context"
	"log"
	"net/http"

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
	cfg, err := env.ParseAs[config.Config]() //cek
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	db := database.New(context.Background(), cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = database.Migrate(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

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
