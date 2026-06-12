package routes

import (
	"hospital-backend/internal/handler"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterAppointmentRoutes(
	router *gin.RouterGroup,
	appointmentHandler *handler.AppointmentHandler,
) {
	appointmentGroup := router.Group("/appointments")

	appointmentGroup.POST(
		"",
		middleware.RequireRole(
			models.RoleReceptionist,
			models.RoleAdmin,
		),
		appointmentHandler.Create,
	)

	appointmentGroup.GET(
		"/patient/:patientId",
		appointmentHandler.GetByPatient,
	)

	appointmentGroup.GET(
		"/doctor/:doctorId",
		appointmentHandler.GetByDoctor,
	)

	appointmentGroup.PATCH(
		"/:id/status",
		middleware.RequireRole(
			models.RoleDoctor,
			models.RoleReceptionist,
		),
		appointmentHandler.UpdateStatus,
	)
	appointmentGroup.GET(
	"",
	appointmentHandler.List,
)

appointmentGroup.PUT(
    "/:id",
    middleware.RequireRole(
        models.RoleReceptionist,
        models.RoleAdmin,
    ),
    appointmentHandler.Reschedule,
)

}
