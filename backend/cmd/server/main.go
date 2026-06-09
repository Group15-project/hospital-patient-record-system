package main

import (
	"log"

	_ "hospital-backend/docs"

	"hospital-backend/internal/config"
	"hospital-backend/internal/database"
	"hospital-backend/internal/handler"
	"hospital-backend/internal/repository"
	"hospital-backend/internal/routes"
	"hospital-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Hospital Management API
// @version 1.0
// @description Hospital Management System API
// @BasePath /
func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect database
	db, err := database.ConnectMySQL(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	validate := validator.New()

	// Repositories
	userRepo := repository.NewUserRepository(db)
	patientRepo := repository.NewPatientRepository(db)
	vitalRepo := repository.NewVitalRepository(db)
	labRepo := repository.NewLabRepository(db)
	prescriptionRepo := repository.NewPrescriptionRepository(db)
	medicationRepo := repository.NewMedicationRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)
	medicalDocumentRepo := repository.NewMedicalDocumentRepository(db)
	billingRepo := repository.NewBillingRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)
	auditRepo := repository.NewAuditRepository(db)
	
	
	
	// Services
	auditService := service.NewAuditService(

	auditRepo,

)

	authService := service.NewAuthService(
		userRepo,
		auditService,
		cfg.JWTSecretKey,

	)
	patientService := service.NewPatientService(
		patientRepo,
	)
	vitalService := service.NewVitalService(
		vitalRepo,
	)
	labService := service.NewLabService(
		labRepo,
	)
	prescriptionService := service.NewPrescriptionService(
		prescriptionRepo,
	)
	medicationService := service.NewMedicationService(

		medicationRepo,
	)

	appointmentService := service.NewAppointmentService(
		appointmentRepo,
	)
	medicalDocumentService := service.NewMedicalDocumentService(
		medicalDocumentRepo,
	)
	billingService := service.NewBillingService(
		billingRepo,
	)
	dashboardService := service.NewDashboardService(

	dashboardRepo,

)


	// Handlers
	authHandler := handler.NewAuthHandler(
		authService,
	)
	patientHandler := handler.NewPatientHandler(
		patientService,
	)
	vitalHandler := handler.NewVitalHandler(
		vitalService,
	)
	labHandler := handler.NewLabHandler(
		labService,
	)
	prescriptionHandler := handler.NewPrescriptionHandler(
		prescriptionService,
	)
	medicationHandler := handler.NewMedicationHandler(

		medicationService,
	)
	appointmentHandler := handler.NewAppointmentHandler(
		appointmentService,
	)

	medicalDocumentHandler := handler.NewMedicalDocumentHandler(
		medicalDocumentService,
	)
	billingHandler := handler.NewBillingHandler(
		billingService,
		validate,
	)
	dashboardHandler := handler.NewDashboardHandler(

	dashboardService,

)

	// Router
	router := gin.Default()

	// Swagger
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	// Register routes
	routes.RegisterRoutes(
		router,
		cfg,
		authHandler,
		patientHandler,
		vitalHandler,
		labHandler,
		prescriptionHandler,
		medicationHandler,
		appointmentHandler,
		medicalDocumentHandler,
		billingHandler,
		dashboardHandler,
	)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("server running on port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
