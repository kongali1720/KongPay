package provider

import "context"

type ProviderType string

const (
	BankTransfer   ProviderType = "BANK_TRANSFER"
	QRIS           ProviderType = "QRIS"
	VirtualAccount ProviderType = "VIRTUAL_ACCOUNT"
	Card           ProviderType = "CARD"
	Crypto         ProviderType = "CRYPTO"
	ManualFiat     ProviderType = "MANUAL_FIAT"
)

type PaymentProvider interface {
	Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error)
	Confirm(ctx context.Context, txID string) (*PaymentStatus, error)
	Cancel(ctx context.Context, txID string) error
	HandleWebhook(ctx context.Context, payload []byte) (*WebhookEvent, error)
	Name() string
	Type() ProviderType
	IsAvailable(ctx context.Context) bool
}

type PaymentRequest struct {
	Amount      float64                `json:"amount"`
	Currency    string                 `json:"currency"`
	Method      string                 `json:"method"`
	CustomerID  string                 `json:"customer_id"`
	MerchantID  string                 `json:"merchant_id"`
	Description string                 `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type PaymentResponse struct {
	TransactionID  string                 `json:"transaction_id"`
	ProviderTxID   string                 `json:"provider_tx_id"`
	Status         string                 `json:"status"`
	RedirectURL    string                 `json:"redirect_url,omitempty"`
	QRCode         string                 `json:"qr_code,omitempty"`
	VirtualAccount string                 `json:"virtual_account,omitempty"`
	Error          string                 `json:"error,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

type PaymentStatus struct {
	Status    string `json:"status"`
	SettledAt string `json:"settled_at,omitempty"`
	Error     string `json:"error,omitempty"`
}

type WebhookEvent struct {
	TransactionID string                 `json:"transaction_id"`
	ProviderTxID  string                 `json:"provider_tx_id"`
	Status        string                 `json:"status"`
	Data          map[string]interface{} `json:"data"`
}
