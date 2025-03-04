package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/middleware"
	"github.com/kedarnacha/gatxel-go/repository"
	"gorm.io/gorm"
)

func SetupAppoinmentRouter(r *gin.Engine, db *gorm.DB) {
	appoinmentRepository := repository.NewAppoinmentRepository(db)
	appoinmentHandler := handler.NewAppoinmentHandler(appoinmentRepository)

	appoinment := r.Group("/appoinment")
	appoinment.Use(middleware.AuthProtected(db))
	{
		appoinment.GET("", appoinmentHandler.GetAllAppoinment)
		appoinment.POST("", middleware.RoleRequired("admin"), appoinmentHandler.CreateAppointment)
		appoinment.GET("/:id", appoinmentHandler.GetAppointmentByID)
		appoinment.PUT("/:id", middleware.RoleRequired("admin"), appoinmentHandler.UpdateAppointmentByID)
		appoinment.DELETE("/:id", middleware.RoleRequired("admin"), appoinmentHandler.DeleteAppoinmentByID)
	}
}
