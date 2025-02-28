package handler

import (
	"gatxel-appointment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service models.AuthService
}

func NewAuthHandler(service models.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var credentials models.AuthCredentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := h.service.Login(ctx, &credentials)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, newUser, err := h.service.Register(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": newUser})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if err := h.service.Logout(ctx, token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
