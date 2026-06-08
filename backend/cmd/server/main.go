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

	// Repositories
	userRepo := repository.NewUserRepository(db)
	patientRepo := repository.NewPatientRepository(db)
	vitalRepo := repository.NewVitalRepository(db)
	labRepo := repository.NewLabRepository(db)
	prescriptionRepo := repository.NewPrescriptionRepository(db)
	medicationRepo := repository.NewMedicationRepository(db)

	// Services
	authService := service.NewAuthService(
		userRepo,
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