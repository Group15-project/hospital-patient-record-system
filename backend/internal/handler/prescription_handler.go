package handler

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



type PrescriptionHandler struct {

	service  service.PrescriptionService
	validate *validator.Validate

}

func NewPrescriptionHandler(

	service service.PrescriptionService,

) *PrescriptionHandler {

	return &PrescriptionHandler{
		service: service,
		validate: validator.New(),
	}

}

func (h *PrescriptionHandler) Create(
	c *gin.Context,
) {

	var req dto.CreatePrescriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request",
			err.Error(),
		)
		return
	}

	doctorID := c.GetUint("user_id")

	err := h.service.Create(
		req,
		doctorID,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to create prescription",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"prescription created",
		nil,
	)
}

func (h *PrescriptionHandler) GetByPatient(
	c *gin.Context,
) {

	id, err := strconv.ParseUint(
		c.Param("patientId"),
		10,
		64,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid patient id",
			nil,
		)
		return
	}

	prescriptions, err := h.service.GetByPatient(
		uint(id),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch prescriptions",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"prescriptions retrieved",
		prescriptions,
	)
}