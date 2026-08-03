package models

import (
	"time"

	"github.com/google/uuid"
)

type Settlement struct {
	ID uuid.UUID `json:"id"`

	BatchReference string `json:"batch_reference"`

	Currency string `json:"currency"`

	TotalAmount float64 `json:"total_amount"`

	TransactionCount int `json:"transaction_count"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`

	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

const (
	SettlementCreated    = "CREATED"
	SettlementProcessing = "PROCESSING"
	SettlementCompleted  = "COMPLETED"
	SettlementFailed     = "FAILED"
)
