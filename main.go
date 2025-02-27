package main

import (
	// "log"

	"context"
	"log"
	"net/http"

	"github.com/YusronAbi/gatxel-go-master/config"
	"github.com/YusronAbi/gatxel-go-master/database"
	"github.com/YusronAbi/gatxel-go-master/repository"
	"github.com/YusronAbi/gatxel-go-master/router"
	"github.com/YusronAbi/gatxel-go-master/service"
	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	// "github.com/joho/godotenv"
	// "gatxel-appointment/config"
	// "gatxel-appointment/routes"
)

var DB *pgxpool.Pool

func main() {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		log.Fatal(err.Error())
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
	router.SetupAppointmentRouter(r, db)
	router.SetupUserRouter(r, db)
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	router.SetupAuthRouter(r, authService.(*service.AuthService))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "server is run")
	})
	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
