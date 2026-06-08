package middleware

import (
	"net/http"

	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {

		roleValue, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "role not found",
			})
			return
		}

		role := roleValue.(string)

		if role == models.RoleSuperAdmin {
			c.Next()
			return
		}

		for _, allowedRole := range roles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
	}
}