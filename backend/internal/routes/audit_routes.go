package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterAuditRoutes(
	router *gin.RouterGroup,
	auditHandler *handler.AuditHandler,
) {
	auditGroup := router.Group("/audit")

	auditGroup.GET(
		"",
		middleware.RequireRole(
			models.RoleAdmin,
			models.RoleSuperAdmin,
		),
		auditHandler.List,
	)
}
