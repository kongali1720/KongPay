package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type LedgerHandler struct {
	DB   *pgx.Conn
	Repo *repositories.LedgerRepository
}

func NewLedgerHandler(
	db *pgx.Conn,
	repo *repositories.LedgerRepository,
) *LedgerHandler {

	return &LedgerHandler{
		DB:   db,
		Repo: repo,
	}
}

func (h *LedgerHandler) ListByWallet(c *gin.Context) {

	walletID, err := uuid.Parse(
		c.Param("wallet_id"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid wallet id",
		})
		return
	}

	entries, err := h.Repo.ListByWalletID(
		c.Request.Context(),
		h.DB,
		walletID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if entries == nil {
		entries = []models.LedgerEntry{}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"entries": entries,
	})
}
