package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// GetWallet returns wallet details
func GetWallet(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "wallet_id": "wallet-123",
        "balance":   1000000,
        "currency":  "IDR",
        "status":    "active",
    })
}

// TopUpWallet handles wallet top-up
func TopUpWallet(c *gin.Context) {
    var req struct {
        Amount float64 `json:"amount"`
        Method string  `json:"method"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Wallet topped up",
        "amount":  req.Amount,
        "new_balance": 1500000,
    })
}
