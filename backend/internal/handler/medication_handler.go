package handler

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicationHandler struct {
	service  service.MedicationService
	validate *validator.Validate
}

func NewMedicationHandler(
	service service.MedicationService,
) *MedicationHandler {
	return &MedicationHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *MedicationHandler) Create(c *gin.Context) {
	var req dto.CreateMedicationRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"validation failed",
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
			"failed to create medication request",
			err.Error(),
		)
		return
	}
	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"medication request created",
		nil,
	)
}

func (h *MedicationHandler) List(c *gin.Context) {
	medications, err := h.service.List()
	
	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"unable to fetch medications",
			err.Error(),
		)
		return
	}
	utils.SuccessResponse(
		c,
		http.StatusOK,
		"medication retrieved",
		medications,
	)
}

func (h *MedicationHandler) AddStock(
	c *gin.Context,
) {
	var req dto.StockAdjustmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"validation failed",
			err.Error(),
		)
		return
	}

	userID := c.GetUint("user_id")

	err := h.service.AddStock(
		req,
		userID,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to add stock",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"stock updated",
		nil,
	)
}