package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kedarnacha/gatxel-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Missing Authorization Header",
			})
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid Token Format",
			})
			ctx.Abort()
			return
		}

		tokenStr := tokenParts[1]
		secret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("Unauthorized - %v", err),
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid Claims",
			})
			ctx.Abort()
			return
		}

		userId, ok := claims["id"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid User ID in Token",
			})
			ctx.Abort()
			return
		}

		var user models.User
		if err := db.Where("id = ?", int64(userId)).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {

			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": "Unauthorized - User Not Found",
			})
			ctx.Abort()
			return
		}
		ctx.Set("userId", userId)
		ctx.Next()
	}
}

func RoleRequired(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Authorization token is missing",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Invalid token",
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Invalid token claims",
			})
			ctx.Abort()
			return
		}

		role, ok := (*claims)["role"].(string)
		if !ok || role == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Role is missing or invalid in token",
			})
			ctx.Abort()
			return
		}

		if !isRoleAllowed(role, allowedRoles) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  "fail",
				"message": "Permission denied",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func isRoleAllowed(role string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}
