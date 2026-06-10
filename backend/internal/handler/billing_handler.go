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

type BillingHandler struct {
	billingService service.BillingService
	validate       *validator.Validate
}

func NewBillingHandler(
	billingService service.BillingService,
	validate *validator.Validate,
) *BillingHandler {
	return &BillingHandler{
		billingService: billingService,
		validate:       validate,
	}
}

func (h *BillingHandler) CreateInvoice(
	c *gin.Context,
) {
	var req dto.CreateInvoiceRequest

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

	if err := h.billingService.CreateInvoice(
		req,
		userID,
	); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"failed to create invoice",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"invoice created successfully",
		nil,
	)
}

func (h *BillingHandler) RecordPayment(
	c *gin.Context,
) {
	var req dto.CreatePaymentRequest

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

	if err := h.billingService.RecordPayment(
		req,
		userID,
	); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"failed to record payment",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"payment recorded successfully",
		nil,
	)
}

func (h *BillingHandler) GetPatientInvoices(
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
			err.Error(),
		)
		return
	}

	invoices, err := h.billingService.GetPatientInvoices(
		uint(patientID),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch invoices",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"invoices fetched successfully",
		invoices,
	)
}