package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterPatientRoutes(
	router *gin.RouterGroup,
	patientHandler *handler.PatientHandler,
) {
	patients := router.Group("/patients")

	patients.POST(
		"",
		patientHandler.Create,
	)

	patients.GET(
		"",
		patientHandler.List,
	)

	patients.GET(
		"/:id",
		patientHandler.GetByID,
	)
	patients.DELETE(
		"/:id",
		middleware.RequireRole(
			models.RoleAdmin,
			models.RoleSuperAdmin,
		),
		patientHandler.Delete,
	)
}