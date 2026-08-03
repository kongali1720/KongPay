package models

import (
	"time"

	"github.com/google/uuid"
)

type SettlementReconciliation struct {
	ID uuid.UUID `json:"id"`

	SettlementID uuid.UUID `json:"settlement_id"`

	TransactionID *uuid.UUID `json:"transaction_id,omitempty"`

	ExpectedAmount float64 `json:"expected_amount"`

	ActualAmount float64 `json:"actual_amount"`

	Difference float64 `json:"difference"`

	Status string `json:"status"`

	Notes string `json:"notes,omitempty"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}
