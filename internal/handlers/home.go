package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kongali1720/KongPay/internal/metadata"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Welcome to KongPay API",
		"data": gin.H{
			"project": "KongPay",
			"status":  "running",
			"version": metadata.Version(),
			"build":   metadata.Build(),
		},
	})
}
