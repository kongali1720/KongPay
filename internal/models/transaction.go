package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID               uuid.UUID `json:"id"`
	ReferenceNo      string    `json:"reference_no"`
	SenderWalletID   uuid.UUID `json:"sender_wallet_id"`
	ReceiverWalletID uuid.UUID `json:"receiver_wallet_id"`
	Amount           float64   `json:"amount"`
	Currency         string    `json:"currency"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
