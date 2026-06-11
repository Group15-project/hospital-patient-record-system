package main

import (
	"log"
	"time"

	_ "hospital-backend/docs"

	"hospital-backend/internal/config"
	"hospital-backend/internal/database"
	"hospital-backend/internal/handler"
	"hospital-backend/internal/repository"
	"hospital-backend/internal/routes"
	"hospital-backend/internal/service"

	"github.com/gin-contrib/cors"
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
	authRepo := repository.NewAuthRepository(db)
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
	medicalRecordRepo :=
		repository.NewMedicalRecordRepository(
			db,
		)
	consultationRepo := repository.NewConsultationRepository(db)

	// Services
	auditService := service.NewAuditService(

		auditRepo,
	)

	authService := service.NewAuthService(
		authRepo,
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

	medicalRecordService :=
		service.NewMedicalRecordService(
			medicalRecordRepo,
		)
	consultationService := service.NewConsultationService(
		consultationRepo,
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
	auditHandler := handler.NewAuditHandler(
		auditService,
	)

	medicalRecordHandler :=
		handler.NewMedicalRecordHandler(
			medicalRecordService,
		)
	consultationHandler := handler.NewConsultationHandler(
		consultationService,
	)

	// Router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
    AllowOrigins: []string{
        "http://127.0.0.1:5500",
        "http://127.0.0.1:5501",
        "http://localhost:5500",
        "http://localhost:5173",

        "https://hospital-patient-record-system-1-omgx.onrender.com",
    },
    AllowMethods: []string{
        "GET",
        "POST",
        "PUT",
        "PATCH",
        "DELETE",
        "OPTIONS",
    },
    AllowHeaders: []string{
        "Origin",
        "Content-Type",
        "Authorization",
    },
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
}))

	// Swagger
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	router.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "ok",
    })
})

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
		auditHandler,
		medicalRecordHandler,
		consultationHandler,
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
