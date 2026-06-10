package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterBIllingRoutes(
	router *gin.RouterGroup,
	billingHandler *handler.BillingHandler,
) {
	billingGroup := router.Group("/billing")

	billingGroup.POST(
		"/invoice",
		middleware.RequireRole(
			models.RoleReceptionist,
			models.RoleBilling,
		),
		billingHandler.CreateInvoice,
	)

	billingGroup.POST(
		"/payment",
		middleware.RequireRole(
			models.RoleReceptionist,
			models.RoleBilling,
		),
		billingHandler.RecordPayment,
	)

	billingGroup.GET(
		"/patient/:patientId",
		billingHandler.GetPatientInvoices,
	)
}
