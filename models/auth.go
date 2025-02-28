package models

import "context"

type AuthCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AuthRepository interface {
	RegisterUser(ctx context.Context, User *User) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

type AuthService interface {
	Login(ctx context.Context, login *AuthCredentials) (string, *User, error)
	Register(ctx context.Context, register *User) (string, *User, error)
	Logout(ctx context.Context, token string) error
}
