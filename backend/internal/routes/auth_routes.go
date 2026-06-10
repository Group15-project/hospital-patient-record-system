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

	authGroup.POST(
		"/login",
		authHandler.Login,
	)
	authGroup.POST(

		"/users",
		middleware.RequireRole(
			models.RoleReceptionist,
			models.RoleAdmin,
		),
		authHandler.RegisterUser,
	)
	authGroup.GET(
	"/doctors",
	authHandler.GetDoctors,
)
authGroup.GET(
	"/profile",

	authHandler.Profile,
)
}