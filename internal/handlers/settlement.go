package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kongali1720/KongPay/internal/settlement"
)

type SettlementHandler struct {
	service *settlement.Service
}

func NewSettlementHandler(
	service *settlement.Service,
) *SettlementHandler {

	return &SettlementHandler{
		service: service,
	}
}

func (h *SettlementHandler) Create(c *gin.Context) {

	var req settlement.CreateSettlementRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	result, err := h.service.CreateBatch(
		c.Request.Context(),
		req.Currency,
		req.Amount,
		req.Count,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":    true,
		"settlement": result,
	})
}
