package handler

import (
	"net/http"

	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuditHandler struct {
	service service.AuditService
}

func NewAuditHandler(
	service service.AuditService,
) *AuditHandler {
	return &AuditHandler{
		service: service,
	}
}

func (h *AuditHandler) List(
	c *gin.Context,
) {

	logs, err := h.service.List(
		1,
		100,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch logs",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"audit logs",
		logs,
	)
}