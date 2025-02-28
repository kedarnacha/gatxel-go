package router

import (
	"github.com/kedarnacha/gatxel-go/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kedarnacha/gatxel-go/handler"
)

func SetupUserRouter(r *gin.Engine, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)
	userHandler := handler.NewHandlerUser(userRepository)
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUserByID)
		userGroup.DELETE("/:id", userHandler.DeleteUserByID)
	}
}
