package router

import (
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/middleware"
	"github.com/kedarnacha/gatxel-go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRouter(r *gin.Engine, authService *service.AuthService, db *gorm.DB) {
	authHandler := handler.NewAuthHandler(authService)

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.POST("/logout", middleware.AuthProtected(db), authHandler.Logout)
}
