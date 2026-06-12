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


type PatientHandler struct {

	service  service.PatientService
	validate *validator.Validate

}

func NewPatientHandler(

	service service.PatientService,

) *PatientHandler {

	return &PatientHandler{
		service: service,
		validate: validator.New(),
	}

}

func (h *PatientHandler) Create(c *gin.Context) {

	var req dto.CreatePatientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid payload",
			err.Error(),
		)
		return
	}

	userID := c.GetUint("user_id")

	patient, err := h.service.Create(
		req,
		userID,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"unable to create patient",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"patient created",
		patient,
	)
}

func (h *PatientHandler) GetByID(c *gin.Context) {

	id, _ := strconv.Atoi(
		c.Param("id"),
	)

	patient, err := h.service.GetByID(
		uint(id),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusNotFound,
			"patient not found",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"patient retrieved",
		patient,
	)
}

func (h *PatientHandler) List(c *gin.Context) {

	patients, err := h.service.List(
		1,
		50,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"unable to fetch patients",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"patients retrieved",
		patients,
	)
}

func (h *PatientHandler) Delete(
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
				"message": "invalid patient id",
			},
		)
		return
	}

	err = h.service.Delete(
		uint(id),
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
			"message": "patient deleted successfully",
		},
	)
}