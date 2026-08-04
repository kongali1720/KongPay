package provider

import (
    "context"
)

// ProviderType defines the type of payment provider
type ProviderType string

const (
    BankTransfer   ProviderType = "BANK_TRANSFER"
    QRIS           ProviderType = "QRIS"
    VirtualAccount ProviderType = "VIRTUAL_ACCOUNT"
    Card           ProviderType = "CARD"
    Crypto         ProviderType = "CRYPTO"
)

// PaymentProvider interface for all payment providers
type PaymentProvider interface {
    // Process initiates a payment
    Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error)
    
    // Confirm checks the status of a payment
    Confirm(ctx context.Context, txID string) (*PaymentStatus, error)
    
    // Cancel cancels a pending payment
    Cancel(ctx context.Context, txID string) error
    
    // HandleWebhook processes incoming webhook from provider
    HandleWebhook(ctx context.Context, payload []byte) (*WebhookEvent, error)
    
    // Name returns the provider name
    Name() string
    
    // Type returns the provider type
    Type() ProviderType
    
    // IsAvailable checks if provider is healthy
    IsAvailable(ctx context.Context) bool
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
    Amount      float64                `json:"amount"`
    Currency    string                 `json:"currency"`
    Method      string                 `json:"method"`
    CustomerID  string                 `json:"customer_id"`
    MerchantID  string                 `json:"merchant_id"`
    Description string                 `json:"description,omitempty"`
    Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// PaymentResponse represents a payment response
type PaymentResponse struct {
    TransactionID  string `json:"transaction_id"`
    ProviderTxID   string `json:"provider_tx_id"`
    Status         string `json:"status"`
    RedirectURL    string `json:"redirect_url,omitempty"`
    QRCode         string `json:"qr_code,omitempty"`
    VirtualAccount string `json:"virtual_account,omitempty"`
    Error          string `json:"error,omitempty"`
}

// PaymentStatus represents payment status
type PaymentStatus struct {
    Status    string `json:"status"`
    SettledAt string `json:"settled_at,omitempty"`
    Error     string `json:"error,omitempty"`
}

// WebhookEvent represents a webhook event
type WebhookEvent struct {
    TransactionID string                 `json:"transaction_id"`
    ProviderTxID  string                 `json:"provider_tx_id"`
    Status        string                 `json:"status"`
    Data          map[string]interface{} `json:"data"`
}
