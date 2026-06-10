package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequirePermission(permission string) gin.HandlerFunc {

	return func(c *gin.Context) {

		permissionsRaw, exists := c.Get("permissions")

		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "permission denied",
			})
			return
		}

		permissions := permissionsRaw.([]string)

		for _, p := range permissions {
			if p == permission {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "permission denied",
		})
	}
}