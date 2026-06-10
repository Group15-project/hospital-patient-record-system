package handler

import (
	"net/http"
	"strconv"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VitalHandler struct {
	service  service.VitalService
	validate *validator.Validate
}

func NewVitalHandler(
	service service.VitalService,
) *VitalHandler {

	return &VitalHandler{
		service: service,
		validate: validator.New(),
	}
}

func (h *VitalHandler) Create(
	c *gin.Context,
) {
	var req dto.CreateVitalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request",
			err.Error(),
		)
		return
	}

	userID := c.GetUint("user_id")

	vital, err := h.service.Create(
		req,
		userID,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to create vital",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"vital recorded",
		vital,
	)
}

func (h *VitalHandler) GetPatientVitals(
	c *gin.Context,
) {

	id, _ := strconv.Atoi(
		c.Param("patientId"),
	)

	vitals, err := h.service.GetPatientVitals(
		uint(id),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch vitals",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"vitals retrieved",
		vitals,
	)
}