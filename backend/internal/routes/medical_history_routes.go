package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterMedicalRecordRoutes(
	router *gin.RouterGroup,
	medicalRecordHandler *handler.MedicalRecordHandler,
) {

	medicalRecords := router.Group(
		"/medical-records",
	)

	medicalRecords.POST(
		"",
		middleware.RequireRole(
			models.RoleDoctor,
			models.RoleAdmin,
			models.RoleSuperAdmin,
		),
		medicalRecordHandler.Create,
	)

	medicalRecords.GET(
		"",
		medicalRecordHandler.List,
	)

	medicalRecords.GET(
		"/:id",
		medicalRecordHandler.GetByID,
	)

	medicalRecords.GET(
		"/patient/:patientId",
		medicalRecordHandler.GetByPatientID,
	)

	medicalRecords.PUT(
		"/:id",
		middleware.RequireRole(
			models.RoleDoctor,
			models.RoleAdmin,
			models.RoleSuperAdmin,
		),
		medicalRecordHandler.Update,
	)

	medicalRecords.DELETE(
		"/:id",
		middleware.RequireRole(
			models.RoleDoctor,
			models.RoleAdmin,
			models.RoleSuperAdmin,
		),
		medicalRecordHandler.Delete,
	)
}