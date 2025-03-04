package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kedarnacha/gatxel-go/helper"
	"github.com/kedarnacha/gatxel-go/models"
	"github.com/kedarnacha/gatxel-go/utils"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repository models.AuthRepository
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) Login(ctx context.Context, login *models.AuthCredentials) (string, *models.User, error) {
	user, err := s.repository.GetUser(ctx, "email = ?", login.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("invalid credentials")
		}
		return "", nil, err
	}

	if !helper.MatchesHash(login.Password, user.Password) {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", nil, fmt.Errorf("JWT secret is not set")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, secret)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthService) Register(ctx context.Context, register *models.User) (string, *models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, fmt.Errorf("failed to hash password: %w", err)
	}
	register.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(ctx, register)
	if err != nil {
		return "", nil, fmt.Errorf("failed to register user: %w", err)
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", nil, fmt.Errorf("JWT secret is not set")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}
	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, secret)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, user, nil
}

func (s *AuthService) register(ctx context.Context, register *models.User) (string, *models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}
	register.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(ctx, register)
	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Second * 86400).Unix(),
	}
	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthService) Logout(ctx context.Context, userID string) error {
	return nil
}
