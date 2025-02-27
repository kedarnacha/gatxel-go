package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupNotificationRouter(r *gin.Engine, db *gorm.DB) {
	notificationRepository := repository.NewNotificationRepository(db)
	productHandler := handler.NewHandlerProduct(notificationRepository)

	notification := r.Group("/notification")
	notification.Use(middleware.AuthProtected(db))
	{
		notification.GET("", productHandler.GetAllProduct)
		notification.POST("", middleware.RoleRequired("admin"), productHandler.CreateProduct)
		notification.GET("/:id", productHandler.GetProductByID)
		notification.PUT("/:id", middleware.RoleRequired("admin"), productHandler.UpdateProductByID)
		notification.DELETE("/:id", middleware.RoleRequired("admin"), productHandler.DeleteProductByID)
	}
}
