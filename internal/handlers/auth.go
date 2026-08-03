package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kongali1720/KongPay/internal/services"
)

type AuthHandler struct {
	service *services.Service
}

func NewAuthHandler(service *services.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req services.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user": user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "login coming soon",
	})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "profile coming soon",
	})
}
