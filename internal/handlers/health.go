package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// HealthCheck returns service health status
func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "service":   "KongPay",
        "version":   "1.0.0-alpha.8.1",
        "status":    "healthy",
        "timestamp": time.Now().UTC().Format(time.RFC3339),
    })
}

// SettlementStats returns settlement statistics
func SettlementStats(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "total":     0,
        "pending":   0,
        "completed": 0,
        "failed":    0,
    })
}

// SettlementStatus returns settlement status for a transaction
func SettlementStatus(c *gin.Context) {
    txID := c.Param("transaction_id")
    c.JSON(http.StatusOK, gin.H{
        "transaction_id": txID,
        "status":         "PENDING",
        "message":        "Settlement status check",
    })
}
