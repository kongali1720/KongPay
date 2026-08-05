package handlers

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/kongali1720/KongPay/internal/services"
)

type CryptoPaymentHandler struct {
    txService *services.TransactionService
}

func NewCryptoPaymentHandler(txService *services.TransactionService) (*CryptoPaymentHandler, error) {
    log.Println("✅ Crypto handler initialized")
    return &CryptoPaymentHandler{
        txService: txService,
    }, nil
}

// GenerateCryptoWallet - Gin handler
func (h *CryptoPaymentHandler) GenerateCryptoWallet(c *gin.Context) {
    var req struct {
        TransactionID string `json:"transaction_id"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Generate dummy address
    address := fmt.Sprintf("0x%x", time.Now().UnixNano())
    privateKey := fmt.Sprintf("0x%x", time.Now().UnixNano()+1000)

    log.Printf("✅ Wallet generated for TX: %s", req.TransactionID)

    c.JSON(http.StatusOK, gin.H{
        "address":        address,
        "private_key":    privateKey,
        "network":        "ethereum",
        "currency":       "ETH",
        "transaction_id": req.TransactionID,
        "message":        "Send ETH to this address to complete payment",
    })
}
