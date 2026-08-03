package settlement

import "github.com/kongali1720/KongPay/internal/models"

type CreateSettlementResponse struct {
	Success    bool               `json:"success"`
	Settlement *models.Settlement `json:"settlement,omitempty"`
	Message    string             `json:"message,omitempty"`
}
