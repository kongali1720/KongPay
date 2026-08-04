package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "service":   "KongPay",
        "version":   "1.0.0-alpha.8.1",
        "status":    "healthy",
        "timestamp": time.Now().UTC().Format(time.RFC3339),
    })
}
