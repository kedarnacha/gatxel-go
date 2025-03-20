package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kedarnacha/gatxel-go/helper"
	"github.com/kedarnacha/gatxel-go/models"
	"github.com/kedarnacha/gatxel-go/utils"

	"github.com/golang-jwt/jwt/v5"
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

func (s *AuthService) Register(ctx context.Context, register *models.User) (string, *models.User, error) {
	if register.Email == "" || register.Password == "" {
		return "", nil, errors.New("email dan password tidak boleh kosong")
	}

	fmt.Println("Mendaftarkan user dengan email:", register.Email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Gagal hashing password:", err)
		return "", nil, errors.New("gagal hashing password")
	}
	register.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(ctx, register)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return "", nil, errors.New("email sudah terdaftar")
		}
		fmt.Println("Error saat menyimpan user:", err)
		return "", nil, errors.New("gagal menyimpan user ke database")
	}

	fmt.Println("User berhasil didaftarkan dengan ID:", user.ID)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("JWT_SECRET belum di-set di env")
		return "", nil, errors.New("server error: JWT_SECRET tidak ditemukan")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Second * 86400).Unix(),
	}
	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, secret)
	if err != nil {
		fmt.Println("Error saat membuat JWT:", err)
		return "", nil, errors.New("gagal membuat token JWT")
	}

	fmt.Println("JWT berhasil dibuat untuk user:", user.Email)
	return token, user, nil
}

func (s *AuthService) Logout(ctx context.Context, userID string) error {
	return nil
}
