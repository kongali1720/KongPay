package settlement

type CreateSettlementRequest struct {
	Currency string  `json:"currency" binding:"required"`
	Amount   float64 `json:"amount" binding:"required,gt=0"`
	Count    int     `json:"transaction_count" binding:"required,gt=0"`
}
