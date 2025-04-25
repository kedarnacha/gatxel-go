package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/middleware"
	"github.com/kedarnacha/gatxel-go/repository"
	"gorm.io/gorm"
)

func SetupAppoinmentDayRouter(r *gin.Engine, db *gorm.DB) {
	appoinmentDayRepository := repository.NewAppoinmentDayRepository(db)
	appoinmentDayHandler := handler.NewAppoinmentDayHandler(appoinmentDayRepository)

	appoinmentDay := r.Group("/appoinment_day")
	appoinmentDay.Use(middleware.AuthProtected(db))
	{
		appoinmentDay.GET("", appoinmentDayHandler.GetAllAppoinmentDay)
		appoinmentDay.POST("", middleware.RoleRequired("admin"), appoinmentDayHandler.CreateAppoinmentDay)
		appoinmentDay.GET("/:id", appoinmentDayHandler.GetAppoinmentDayByID)
		appoinmentDay.PUT("/:id", middleware.RoleRequired("admin"), appoinmentDayHandler.UpdateAppoinmentDayByID)
		appoinmentDay.DELETE("/:id", middleware.RoleRequired("admin"), appoinmentDayHandler.DeleteAppoinmentDayByID)
	}
}
