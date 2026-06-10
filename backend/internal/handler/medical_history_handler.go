package handler

import (
	"log"
	"net/http"
	"strconv"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicalRecordHandler struct {
	service  service.MedicalRecordService
	validate *validator.Validate
}

func NewMedicalRecordHandler(
	service service.MedicalRecordService,
) *MedicalRecordHandler {
	return &MedicalRecordHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *MedicalRecordHandler) Create(
	c *gin.Context,
) {

	var req dto.CreateMedicalRecordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		 log.Println("MEDICAL RECORD ERROR:", err)
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid payload",
			err.Error(),
		)
		return
	}

	userID := c.GetUint("user_id")

	record, err := h.service.Create(
		req,
		userID,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to create medical record",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"medical record created",
		record,
	)
}

func (h *MedicalRecordHandler) List(
	c *gin.Context,
) {

	records, err := h.service.List()

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch medical records",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"medical records retrieved",
		records,
	)
}

func (h *MedicalRecordHandler) GetByID(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid id",
			nil,
		)
		return
	}

	record, err := h.service.GetByID(
		uint(id),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusNotFound,
			"medical record not found",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"medical record retrieved",
		record,
	)
}

func (h *MedicalRecordHandler) GetByPatientID(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("patientId"),
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

	records, err := h.service.GetByPatientID(
		uint(id),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch patient records",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"patient records retrieved",
		records,
	)
}

func (h *MedicalRecordHandler) Update(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid id",
			nil,
		)
		return
	}

	var req dto.UpdateMedicalRecordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid payload",
			err.Error(),
		)
		return
	}

	record, err := h.service.Update(
		uint(id),
		req,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to update medical record",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"medical record updated",
		record,
	)
}

func (h *MedicalRecordHandler) Delete(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid id",
			nil,
		)
		return
	}

	err = h.service.Delete(
		uint(id),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to delete medical record",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"medical record deleted",
		nil,
	)
}