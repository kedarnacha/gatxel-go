package router

import (
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/service"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine, authService *service.AuthService) {

	authHandler := handler.NewAuthHandler(authService)
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)

	}

}
