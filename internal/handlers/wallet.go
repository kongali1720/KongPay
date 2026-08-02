package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/services"
)

type WalletHandler struct {
	Service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
	return &WalletHandler{
		Service: service,
	}
}

type CreateWalletRequest struct {
	UserID   string `json:"user_id"`
	Currency string `json:"currency"`
}

func (h *WalletHandler) CreateWallet(c *gin.Context) {

	var req CreateWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id",
		})
		return
	}

	wallet, err := h.Service.CreateWallet(
		c.Request.Context(),
		userID,
		req.Currency,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

func (h *WalletHandler) GetWallet(c *gin.Context) {

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "GetWallet endpoint",
		"id":      id,
	})
}
