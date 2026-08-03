package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ReferenceNo string    `json:"reference_no" db:"reference_no"`

	SenderWalletID   uuid.UUID `json:"sender_wallet_id" db:"sender_wallet_id"`
	ReceiverWalletID uuid.UUID `json:"receiver_wallet_id" db:"receiver_wallet_id"`

	Amount   float64 `json:"amount" db:"amount"`
	Currency string  `json:"currency" db:"currency"`

	Status string `json:"status" db:"status"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
