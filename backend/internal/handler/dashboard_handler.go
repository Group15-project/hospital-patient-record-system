package handler

import (
	"net/http"

	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service service.DashboardService
}

func NewDashboardHandler(
	service service.DashboardService,
) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

func (h *DashboardHandler) Summary(
	c *gin.Context,
) {

	data, err := h.service.GetSummary()

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to load dashboard",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"dashboard summary",
		data,
	)
}

func (h *DashboardHandler) LowStock(
	c *gin.Context,
) {

	data, err := h.service.
		GetLowStockMedications()

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to load inventory alerts",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"inventory alerts",
		data,
	)
}

func (h *DashboardHandler) Today(
	c *gin.Context,
) {

	data, err := h.service.GetTodaySummary()

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to load today's dashboard",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"today dashboard",
		data,
	)
}