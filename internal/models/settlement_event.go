package models

import (
	"time"

	"github.com/google/uuid"
)

type SettlementEvent struct {
	ID uuid.UUID `json:"id"`

	SettlementID uuid.UUID `json:"settlement_id"`

	EventType string `json:"event_type"`

	OldStatus string `json:"old_status"`

	NewStatus string `json:"new_status"`

	CreatedAt time.Time `json:"created_at"`
}

const (
	SettlementEventCreated = "SETTLEMENT_CREATED"

	SettlementEventProcessing = "SETTLEMENT_PROCESSING"

	SettlementEventCompleted = "SETTLEMENT_COMPLETED"

	SettlementEventFailed = "SETTLEMENT_FAILED"
)
