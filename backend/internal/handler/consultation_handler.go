package handler

import (
	"net/http"
	"strconv"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ConsultationHandler struct {
	service service.ConsultationService
}

func NewConsultationHandler(
	service service.ConsultationService,
) *ConsultationHandler {
	return &ConsultationHandler{
		service: service,
	}
}

func (h *ConsultationHandler) Create(
	c *gin.Context,
) {

	var req dto.CreateConsultationRequest

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

	doctorID := c.GetUint("user_id")

	consultation, err :=
		h.service.CreateConsultation(
			req,
			doctorID,
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
		http.StatusCreated,
		gin.H{
			"success": true,
			"data": consultation,
		},
	)
}

func (h *ConsultationHandler) AddDiagnosis(
	c *gin.Context,
) {

	var req dto.AddDiagnosisRequest

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

	doctorID := c.GetUint("user_id")

	err := h.service.AddDiagnosis(
		req,
		doctorID,
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
		http.StatusCreated,
		gin.H{
			"success": true,
			"message": "diagnosis added successfully",
		},
	)
}

func (h *ConsultationHandler) GetByID(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid consultation id",
			},
		)

		return
	}

	consultation, err :=
		h.service.GetByID(
			uint(id),
		)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
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
			"data": consultation,
		},
	)
}

func (h *ConsultationHandler) GetDiagnoses(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "invalid consultation id",
			},
		)

		return
	}

	diagnoses, err :=
		h.service.GetDiagnoses(
			uint(id),
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
			"data": diagnoses,
		},
	)
}

func (h *ConsultationHandler) GetByPatient(
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

	consultations, err :=
		h.service.GetByPatient(
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
			"data": consultations,
		},
	)
}