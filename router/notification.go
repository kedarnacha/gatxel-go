package router

import (
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/middleware"
	"github.com/kedarnacha/gatxel-go/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupNotificationRouter(r *gin.Engine, db *gorm.DB) {
	notificationRepository := repository.NewNotificationRepository(db)
	notificationHandler := handler.NewNotificationHandler(notificationRepository)

	notification := r.Group("/notification")
	notification.Use(middleware.AuthProtected(db))
	{
		notification.GET("", notificationHandler.GetAllNotifications)
		notification.POST("", middleware.RoleRequired("admin"), notificationHandler.CreateNotification)
		notification.GET("/:id", notificationHandler.GetNotificationByID)
		notification.PUT("/:id", middleware.RoleRequired("admin"), notificationHandler.UpdateNotificationByID)
		notification.DELETE("/:id", middleware.RoleRequired("admin"), notificationHandler.DeleteNotificationByID)
	}
}
