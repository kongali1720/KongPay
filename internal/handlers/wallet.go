package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateWalletRequest struct {
	UserID   string `json:"user_id"`
	Currency string `json:"currency"`
}

func CreateWallet(c *gin.Context) {

	var req CreateWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{

		"message": "Wallet endpoint ready",

		"user_id": req.UserID,

		"currency": req.Currency,

		"balance": 0,

		"status": "ACTIVE",
	})
}
