package router

import (
	"gatxel-appointment/handler"
	"gatxel-appointment/service"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine, authService *service.AuthService) {

	authHandler := handler.NewAuthHandler(authService)

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.POST("/logout", authHandler.Logout)
}
