package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterVitalRoutes(
	router *gin.RouterGroup,
	vitalHandler *handler.VitalHandler,
) {
vitalGroup :=router.Group("/vitals")

	vitalGroup.POST(
		"",
		middleware.RequireRole(
			models.RoleNurse,
			models.RoleDoctor,
		),
		vitalHandler.Create,
	)

	vitalGroup.GET(
		"/patient/:patientId",
		vitalHandler.GetPatientVitals,
	)

}