package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterLabRoutes(
	router *gin.RouterGroup,
	labHandler *handler.LabHandler,
) {
	labGroup := router.Group("/labs")

	labGroup.POST(
		"/request",
		middleware.RequireRole(models.RoleDoctor),
		labHandler.CreateRequest,
	)
	labGroup.POST(
		"/result",
		middleware.RequireRole(models.RoleLabTechnician),
		labHandler.UploadResult,
	)
	labGroup.GET(
		"/patient/:patientId/requests",
		labHandler.GetPatientRequests,
	)
	labGroup.GET(
		"/patient/:patientId/results",
		labHandler.GetPatientResults,
	)
}
