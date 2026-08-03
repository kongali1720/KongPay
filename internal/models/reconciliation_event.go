package models

import (
	"time"

	"github.com/google/uuid"
)

type ReconciliationEvent struct {
	ID uuid.UUID `json:"id"`

	SettlementID uuid.UUID `json:"settlement_id"`

	ReconciliationID uuid.UUID `json:"reconciliation_id"`

	EventType string `json:"event_type"`

	OldStatus string `json:"old_status"`

	NewStatus string `json:"new_status"`

	CreatedAt time.Time `json:"created_at"`
}
