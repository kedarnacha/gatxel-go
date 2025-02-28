package router

import (
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/service"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine, authService *service.AuthService) {

	authHandler := handler.NewAuthHandler(authService)

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.POST("/logout", authHandler.Logout)
}
