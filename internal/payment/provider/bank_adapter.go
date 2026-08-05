package provider

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
)

type BankAdapter struct {
    apiKey  string
    baseURL string
}

func NewBankAdapter(apiKey, baseURL string) *BankAdapter {
    return &BankAdapter{
        apiKey:  apiKey,
        baseURL: baseURL,
    }
}

func (b *BankAdapter) Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    return &PaymentResponse{
        TransactionID: fmt.Sprintf("KONG-%d", time.Now().UnixNano()),
        ProviderTxID:  fmt.Sprintf("BANK-%d", time.Now().UnixNano()),
        Status:        "PENDING",
        RedirectURL:   b.baseURL + "/payment/" + fmt.Sprintf("%d", time.Now().UnixNano()),
    }, nil
}

func (b *BankAdapter) Confirm(ctx context.Context, txID string) (*PaymentStatus, error) {
    return &PaymentStatus{
        Status:    "SUCCESS",
        SettledAt: time.Now().UTC().Format(time.RFC3339),
    }, nil
}

func (b *BankAdapter) Cancel(ctx context.Context, txID string) error {
    return nil
}

func (b *BankAdapter) HandleWebhook(ctx context.Context, payload []byte) (*WebhookEvent, error) {
    var data map[string]interface{}
    if err := json.Unmarshal(payload, &data); err != nil {
        return nil, fmt.Errorf("invalid webhook payload: %w", err)
    }
    return &WebhookEvent{
        TransactionID: getStringValue(data, "transaction_id"),
        ProviderTxID:  getStringValue(data, "provider_tx_id"),
        Status:        getStringValue(data, "status"),
        Data:          data,
    }, nil
}

func (b *BankAdapter) Name() string {
    return "Bank Transfer Provider"
}

func (b *BankAdapter) Type() ProviderType {
    return BankTransfer
}

func (b *BankAdapter) IsAvailable(ctx context.Context) bool {
    return true
}
