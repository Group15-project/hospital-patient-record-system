package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterConsultationRoutes(
	router *gin.RouterGroup,
	consultationHandler *handler.ConsultationHandler,
) {

	consultations := router.Group(
		"/consultations",
	)

	consultations.POST(
		"",
		middleware.RequireRole(
			models.RoleDoctor,
		),
		consultationHandler.Create,
	)

	consultations.POST(
		"/diagnosis",
		middleware.RequireRole(
			models.RoleDoctor,
		),
		consultationHandler.AddDiagnosis,
	)

	consultations.GET(
		"/:id",
		consultationHandler.GetByID,
	)

	consultations.GET(
		"/:id/diagnoses",
		consultationHandler.GetDiagnoses,
	)
	consultations.GET(
	"/patient/:patientId",
	consultationHandler.GetByPatient,
)
}