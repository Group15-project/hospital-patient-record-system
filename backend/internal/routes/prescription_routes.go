package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterPrescriptionRoutes(
	router *gin.RouterGroup,
	prescriptionHandler *handler.PrescriptionHandler,
) {
	prescriptionGroup := router.Group("/prescriptions")

	prescriptionGroup.POST(

		"",
		middleware.RequireRole(
			models.RoleDoctor,
		),
		prescriptionHandler.Create,
	)
	prescriptionGroup.GET(
		"/patient/:patientId",
		prescriptionHandler.GetByPatient,
	)
}
