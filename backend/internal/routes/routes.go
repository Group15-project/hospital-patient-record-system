package routes

import (
	"hospital-backend/internal/config"
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	patientHandler *handler.PatientHandler,
	vitalHandler *handler.VitalHandler,
	labHandler *handler.LabHandler,
	prescriptionHandler *handler.PrescriptionHandler,
	medicationHandler *handler.MedicationHandler,
) {
	api := router.Group("/api/v1")

	RegisterAuthRoutes(api, authHandler)

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))

	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "authenticated",
		})
	})

	RegisterPatientRoutes(
		protected,
		patientHandler,
	)

	RegisterVitalRoutes(protected, vitalHandler)
}