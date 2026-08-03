package models

import (
	"time"

	"github.com/google/uuid"
)

type LedgerEntry struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	WalletID      uuid.UUID `json:"wallet_id"`

	EntryType string `json:"entry_type"`

	Amount float64 `json:"amount"`

	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`

	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}
