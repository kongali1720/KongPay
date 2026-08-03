package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type TransactionHandler struct {
	repo *repositories.TransactionRepository
}

func NewTransactionHandler(
	repo *repositories.TransactionRepository,
) *TransactionHandler {

	return &TransactionHandler{
		repo: repo,
	}
}

func (h *TransactionHandler) List(c *gin.Context) {

	transactions, err := h.repo.List(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if transactions == nil {
		transactions = []models.Transaction{}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"transactions": transactions,
	})
}
