package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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

func (h *SettlementHandler) Get(c *gin.Context) {

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	result, err := h.service.FindByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"settlement": result,
	})
}

func (h *SettlementHandler) Process(c *gin.Context) {

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	result, err := h.service.Process(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"settlement": result,
	})
}
