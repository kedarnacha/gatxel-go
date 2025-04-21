package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/middleware"
	"github.com/kedarnacha/gatxel-go/repository"
	"gorm.io/gorm"
)

func SetupAppoinmentSlotRouter(r *gin.Engine, db *gorm.DB) {
	appoinmentSlotRepository := repository.NewAppoinmentSlotRepository(db)
	appoinmentSlotHandler := handler.NewAppoinmentSlotHandler(appoinmentSlotRepository)

	appoinmentSlot := r.Group("/appoinment-slot")
	appoinmentSlot.Use(middleware.AuthProtected(db))
	{
		appoinmentSlot.GET("", appoinmentSlotHandler.GetAllAppoinmentSlot)
		appoinmentSlot.POST("", middleware.RoleRequired("admin"), appoinmentSlotHandler.CreateAppoinmentSlot)
		appoinmentSlot.GET("/:id", appoinmentSlotHandler.GetAppoinmentSlotByID)
		appoinmentSlot.PUT("/:id", middleware.RoleRequired("admin"), appoinmentSlotHandler.UpdateAppoinmentSlotByID)
		appoinmentSlot.DELETE("/:id", middleware.RoleRequired("admin"), appoinmentSlotHandler.DeleteAppoinmentSlotByID)
	}
}
