package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(
	router *gin.RouterGroup,
	dashboardHandler *handler.DashboardHandler,
) {
	dashboardGroup := router.Group("/dashboard")

	dashboardGroup.GET(
		"/summary",
		middleware.RequireRole(
			models.RoleAdmin,
			models.RoleSuperAdmin,
		),
		dashboardHandler.Summary,
	)

	dashboardGroup.GET(
		"/low-stock",
		middleware.RequireRole(
			models.RoleAdmin,
			models.RoleSuperAdmin,
			models.RolePharmacist,
		),
		dashboardHandler.LowStock,
	)

	dashboardGroup.GET(
		"/today",
		middleware.RequireRole(
			models.RoleAdmin,
			models.RoleSuperAdmin,
		),
		dashboardHandler.Today,
	)
}
