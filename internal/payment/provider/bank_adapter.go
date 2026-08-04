package provider

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
)

// BankAdapter implements PaymentProvider for bank transfers
type BankAdapter struct {
    apiKey  string
    baseURL string
}

// NewBankAdapter creates a new bank adapter
func NewBankAdapter(apiKey, baseURL string) *BankAdapter {
    return &BankAdapter{
        apiKey:  apiKey,
        baseURL: baseURL,
    }
}

// Process initiates a bank transfer
func (b *BankAdapter) Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    // TODO: Implement actual bank API call
    // 1. Build request to bank API
    // 2. Send HTTP request
    // 3. Parse response
    
    return &PaymentResponse{
        TransactionID: generateTxID("KONG"),
        ProviderTxID:  generateTxID("BANK"),
        Status:        "PENDING",
        RedirectURL:   b.baseURL + "/payment/" + generateTxID("REF"),
    }, nil
}

// Confirm checks bank transfer status
func (b *BankAdapter) Confirm(ctx context.Context, txID string) (*PaymentStatus, error) {
    // TODO: Implement status check
    return &PaymentStatus{
        Status:    "SUCCESS",
        SettledAt: time.Now().UTC().Format(time.RFC3339),
    }, nil
}

// Cancel cancels a bank transfer
func (b *BankAdapter) Cancel(ctx context.Context, txID string) error {
    // TODO: Implement cancellation
    return nil
}

// HandleWebhook processes bank webhook
func (b *BankAdapter) HandleWebhook(ctx context.Context, payload []byte) (*WebhookEvent, error) {
    var data map[string]interface{}
    if err := json.Unmarshal(payload, &data); err != nil {
        return nil, fmt.Errorf("invalid webhook payload: %w", err)
    }
    
    return &WebhookEvent{
        TransactionID: getString(data, "transaction_id"),
        ProviderTxID:  getString(data, "provider_tx_id"),
        Status:        getString(data, "status"),
        Data:          data,
    }, nil
}

// Name returns provider name
func (b *BankAdapter) Name() string {
    return "Bank Transfer Provider"
}

// Type returns provider type
func (b *BankAdapter) Type() ProviderType {
    return BankTransfer
}

// IsAvailable checks if bank API is healthy
func (b *BankAdapter) IsAvailable(ctx context.Context) bool {
    // TODO: Implement health check
    return true
}

// Helper functions
func generateTxID(prefix string) string {
    return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

func getString(m map[string]interface{}, key string) string {
    if val, ok := m[key]; ok {
        if str, ok := val.(string); ok {
            return str
        }
    }
    return ""
}
