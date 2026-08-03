package models

import (
	"time"

	"github.com/google/uuid"
)

type LedgerEntry struct {
	ID              uuid.UUID `json:"id" db:"id"`

	TransactionID   uuid.UUID `json:"transaction_id" db:"transaction_id"`

	WalletID        uuid.UUID `json:"wallet_id" db:"wallet_id"`

	EntryType       string    `json:"entry_type" db:"entry_type"`

	Amount          float64   `json:"amount" db:"amount"`

	BalanceBefore   float64   `json:"balance_before" db:"balance_before"`

	BalanceAfter    float64   `json:"balance_after" db:"balance_after"`

	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
