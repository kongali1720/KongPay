package payment

import "github.com/google/uuid"

type TransferRequest struct {
	SenderWalletID   uuid.UUID `json:"sender_wallet_id" binding:"required"`
	ReceiverWalletID uuid.UUID `json:"receiver_wallet_id" binding:"required"`
	Amount           float64   `json:"amount" binding:"required,gt=0"`
	Currency         string    `json:"currency"`
}
