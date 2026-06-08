package routes

import (
	"hospital-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(
	router *gin.RouterGroup,
	authHandler *handler.AuthHandler,
) {
	auth := router.Group("/auth")

	auth.POST(
		"/login",
		authHandler.Login,
	)
}