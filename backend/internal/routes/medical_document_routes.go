package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterMedicalDocumentRoutes(
	router *gin.RouterGroup,
	medicalDocumentHandler *handler.MedicalDocumentHandler,
) {
	documentGroup := router.Group("/documents")

	documentGroup.POST(
		"/upload",
		middleware.RequireRole(
			models.RoleDoctor,
			models.RoleNurse,
			models.RoleLabTechnician,
		),
		medicalDocumentHandler.Upload,
	)

	documentGroup.GET(
		"/patient/:patientId",
		medicalDocumentHandler.GetPatientDocuments,
	)
}
