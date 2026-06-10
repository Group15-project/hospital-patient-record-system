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
	appointmentHandler *handler.AppointmentHandler,
	medicalDocumentHandler *handler.MedicalDocumentHandler,
	billingHandler *handler.BillingHandler,
	dashboardHandler *handler.DashboardHandler,
	auditHandler *handler.AuditHandler,
	medicalRecordHandler *handler.MedicalRecordHandler,
	consultationHandler *handler.ConsultationHandler,

) {
	api := router.Group("/api/v1")

	RegisterAuthRoutes(api, authHandler, cfg)

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))

	protected.GET(
	"/profile",
	authHandler.Profile,
)

	RegisterPatientRoutes(
		protected,
		patientHandler,
	)

	RegisterVitalRoutes(protected, vitalHandler)
	RegisterDashboardRoutes(
		protected,
		dashboardHandler,
	)
	RegisterAppointmentRoutes(protected, appointmentHandler)
	RegisterAuditRoutes(protected, auditHandler)
	RegisterBIllingRoutes(protected, billingHandler)
	RegisterLabRoutes(protected, labHandler)
	RegisterMedicalDocumentRoutes(protected, medicalDocumentHandler)
	RegisterMedicationRoutes(protected, medicationHandler)

	RegisterPrescriptionRoutes(protected, prescriptionHandler)
	RegisterMedicalRecordRoutes(
		protected,
		medicalRecordHandler,
	)
	RegisterConsultationRoutes(protected, consultationHandler)
}
