package routes

import (
	"hospital-backend/internal/handler"

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
}