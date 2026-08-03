package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kongali1720/KongPay/internal/payment"
)

type PaymentHandler struct {
	service *payment.Service
}

func NewPaymentHandler(
	service *payment.Service,
) *PaymentHandler {

	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) Transfer(c *gin.Context) {

	var req payment.TransferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	result, err := h.service.Transfer(
		c.Request.Context(),
		req,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
