package routes

import (
	"hospital-backend/internal/config"
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(
	router *gin.RouterGroup,
	authHandler *handler.AuthHandler,
	cfg *config.Config,
) {
	authGroup := router.Group("/auth")

	authGroup.POST("/login", authHandler.Login)

	protected := authGroup.Group("")
	protected.Use(middleware.AuthMiddleware(cfg))

	protected.POST(
		"/users",
		middleware.RequireRole(models.RoleSuperAdmin),
		authHandler.RegisterUser,
	)

	protected.GET(
		"/roles",
		middleware.RequireRole(models.RoleSuperAdmin),
		authHandler.GetRoles,
	)

	protected.DELETE(
		"/users/:id",
		middleware.RequireRole(models.RoleSuperAdmin),
		authHandler.DeleteUser,
	)

	protected.GET(
		"/profile",
		authHandler.Profile,
	)

	protected.GET(
		"/doctors",
		authHandler.GetDoctors,
	)
	protected.GET(
    "/users",
    middleware.RequireRole(models.RoleSuperAdmin),
    authHandler.GetStaff,
)
}
