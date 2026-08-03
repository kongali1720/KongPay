package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/reconciliation"
)

type ReconciliationHandler struct {
	service *reconciliation.Service
}

func NewReconciliationHandler(
	service *reconciliation.Service,
) *ReconciliationHandler {

	return &ReconciliationHandler{
		service: service,
	}
}

type ReconcileRequest struct {
	ExpectedAmount float64 `json:"expected_amount"`
	ActualAmount   float64 `json:"actual_amount"`
}

func (h *ReconciliationHandler) Reconcile(
	c *gin.Context,
) {

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	var req ReconcileRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	result, err := h.service.Reconcile(
		c.Request.Context(),
		id,
		req.ExpectedAmount,
		req.ActualAmount,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"reconciliation": result,
	})
}
