package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kongali1720/KongPay/internal/metadata"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Health check successful",
		"data": gin.H{
			"status":      "ok",
			"service":     "KongPay",
			"version":     metadata.Version(),
			"build":       metadata.Build(),
			"codename":    metadata.Codename(),
			"environment": "development",
			"time":        time.Now().Format(time.RFC3339),
		},
	})
}
