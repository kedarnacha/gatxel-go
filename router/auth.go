package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/service"
)

func SetupAuthRouter(r *gin.Engine, authService *service.AuthService) {

	authHandler := handler.NewAuthHandler(authService)

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.POST("/logout", authHandler.Logout)
}
