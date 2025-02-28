package repository

import (
	"context"
	"github.com/kedarnacha/gatxel-go/models"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUserByID(ctx context.Context, id int64, data map[string]interface{}) (*models.User, error)
	DeleteUserByID(ctx context.Context, id int64) error
}
