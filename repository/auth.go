package repository

import (
	"context"
	"gatxel-appointment/models"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, User *models.User) (*models.User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error)
}

type AuthService interface {
	Login(ctx context.Context, login *models.AuthCredentials) (string, *models.User, error)
	Register(ctx context.Context, register *models.User) (string, *models.User, error)
	Logout(ctx context.Context, token string) error
}
