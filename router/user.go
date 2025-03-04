package router

import (
	"github.com/kedarnacha/gatxel-go/handler"
	"github.com/kedarnacha/gatxel-go/middleware"
	"github.com/kedarnacha/gatxel-go/repository"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) {
	UserRepository := repository.NewUserRepository(db)
	UserHandler := handler.NewUserHandler(UserRepository)

	user := r.Group("/user")
	user.Use(middleware.AuthProtected(db), middleware.RoleRequired("admin"))
	{
		user.GET("", UserHandler.GetAllUser)
		user.GET("/:id", UserHandler.GetUserByID)
		user.PUT("/:id", UserHandler.UpdateUserByID)
		user.DELETE("/:id", UserHandler.DeleteUserByID)
	}
}
