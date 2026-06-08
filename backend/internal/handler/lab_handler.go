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

type LabHandler struct {
	service  service.LabService
	validate *validator.Validate
}

func NewLabHandler(

	service service.LabService,

) *LabHandler {

	return &LabHandler{
		service:  service,
		validate: validator.New(),
	}

}

func (h *LabHandler) CreateRequest(
	c *gin.Context,
) {
	var req dto.CreateLabRequestRequest

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

	doctorID := c.GetUint("user_id")

	err := h.service.CreateRequest(
		req,
		doctorID,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to create lab request",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"lab request created",
		nil,
	)
}

func (h *LabHandler) UploadResult(
	c *gin.Context,
) {
	var req dto.UploadLabResultRequest

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

	labTechID := c.GetUint("user_id")

	err := h.service.UploadResult(
		req,
		labTechID,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to upload result",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"lab result uploaded",
		nil,
	)
}

func (h *LabHandler) GetPatientRequests(
	c *gin.Context,
) {

	patientID, err := strconv.ParseUint(
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

	requests, err := h.service.GetPatientRequests(
		uint(patientID),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch requests",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"lab requests retrieved",
		requests,
	)
}

func (h *LabHandler) GetPatientResults(
	c *gin.Context,
) {

	patientID, err := strconv.ParseUint(
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

	results, err := h.service.GetPatientResults(
		uint(patientID),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch results",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"lab results retrieved",
		results,
	)
}
