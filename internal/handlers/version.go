package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kongali1720/KongPay/internal/metadata"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Version information",
		"data": gin.H{
			"application": "KongPay",
			"version":     metadata.Version(),
			"build":       metadata.Build(),
			"codename":    metadata.Codename(),
		},
	})
}
