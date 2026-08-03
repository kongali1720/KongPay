package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/repositories"
)

type SettlementEventHandler struct {
	repo *repositories.SettlementEventRepository
}

func NewSettlementEventHandler(
	repo *repositories.SettlementEventRepository,
) *SettlementEventHandler {

	return &SettlementEventHandler{
		repo: repo,
	}
}

func (h *SettlementEventHandler) List(c *gin.Context) {

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

	events, err := h.repo.ListBySettlementID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"events":  events,
	})
}
