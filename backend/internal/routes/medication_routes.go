package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterMedicationRoutes(
	router *gin.RouterGroup,
	medicationHandler *handler.MedicationHandler,
) {
	medicationGroup := router.Group("/medications")

	medicationGroup.POST(
		"",
		middleware.RequireRole(
			models.RolePharmacist,
			models.RoleAdmin,
		),
		medicationHandler.Create,
	)

	medicationGroup.GET(
		"",
		medicationHandler.List,
	)

	medicationGroup.POST(
		"/stock",
		middleware.RequireRole(
			models.RolePharmacist,
		),
		medicationHandler.AddStock,
	)
}
