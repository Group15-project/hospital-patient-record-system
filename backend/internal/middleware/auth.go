package middleware

import (
	"net/http"
	"strings"
	"hospital-backend/internal/config"
	"hospital-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authorization header required",
			})
			return
		}
		parts := strings.Split(authHeader, " ")

if len(parts) != 2 || parts[0] != "Bearer" {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "invalid authorization format",
	})
	return
}

		tokenString := parts[1]
		claims, err := utils.ValidateToken(
			tokenString,
			cfg.JWTSecretKey,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}