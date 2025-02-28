package router

import (
	"gatxel-appointment/handler"
	"gatxel-appointment/middleware"
	"gatxel-appointment/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupUserRouter(r *gin.Engine, db *pgxpool.Pool) {
	UserRepository := repository.NewUserRepository(db)
	userHandler := handler.NewHandlerUser(UserRepository)
	user := r.Group("/user")
	user.Use(middleware.AuthProtected(db), middleware.RoleRequired("admin"))

	{
		user.GET("/", userHandler.GetAllUsers)
		user.POST("/", userHandler.CreateUser)
		user.GET("/:id", userHandler.GetUserByID)
		user.PUT("/:id", userHandler.UpdateUserByID)
		user.DELETE("/:id", userHandler.DeleteUserByID)
	}
}
