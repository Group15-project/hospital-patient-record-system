package handler

import (
	"net/http"
	"strconv"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	service service.AppointmentService
}

func NewAppointmentHandler(
	service service.AppointmentService,
) *AppointmentHandler {
	return &AppointmentHandler{
		service: service,
	}
}

func (h *AppointmentHandler) Create(
	c *gin.Context,
) {
	var req dto.CreateAppointmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	userID := c.GetUint("user_id")

	appointment, err := h.service.Create(
		req,
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"success": true,
			"message": "appointment created successfully",
			"data":    appointment,
		},
	)
}

func (h *AppointmentHandler) GetByPatient(
	c *gin.Context,
) {
	patientID, err := strconv.ParseUint(
		c.Param("patientId"),
		10,
		64,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid patient id",
			},
		)
		return
	}

	appointments, err := h.service.GetByPatient(
		uint(patientID),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data":    appointments,
		},
	)
}

func (h *AppointmentHandler) GetByDoctor(
	c *gin.Context,
) {
	doctorID, err := strconv.ParseUint(
		c.Param("doctorId"),
		10,
		64,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid doctor id",
			},
		)
		return
	}

	appointments, err := h.service.GetByDoctor(
		uint(doctorID),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data":    appointments,
		},
	)
}

func (h *AppointmentHandler) UpdateStatus(
	c *gin.Context,
) {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid appointment id",
			},
		)
		return
	}

	var req struct {
		Status models.AppointmentStatus `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	err = h.service.UpdateStatus(
		uint(id),
		req.Status,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "appointment status updated successfully",
		},
	)
}

func (h *AppointmentHandler) List(
	c *gin.Context,
) {

	appointments, err :=
		h.service.List()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data": appointments,
		},
	)
}

func (h *AppointmentHandler) Reschedule(
	c *gin.Context,
) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid appointment id",
			},
		)
		return
	}

	var req dto.RescheduleAppointmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	err = h.service.Reschedule(
		uint(id),
		req.AppointmentDate,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "appointment rescheduled successfully",
		},
	)
}