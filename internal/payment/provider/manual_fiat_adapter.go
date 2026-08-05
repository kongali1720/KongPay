package provider

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
)

type ManualFiatAdapter struct {
    bankName        string
    accountName     string
    accountNumber   string
    instructions    string
}

func NewManualFiatAdapter() *ManualFiatAdapter {
    return &ManualFiatAdapter{
        bankName:      getEnv("MANUAL_FIAT_BANK_NAME", "MANDIRI"),
        accountName:   getEnv("MANUAL_FIAT_ACCOUNT_NAME", "RAZALEE SYAHDAN"),
        accountNumber: getEnv("MANUAL_FIAT_ACCOUNT_NUMBER", "1270014630154"),
        instructions:  getEnv("MANUAL_FIAT_INSTRUCTIONS", "Transfer ke MANDIRI a.n Razalee Syahdan, No Rek: 1270014630154. Konfirmasi via WhatsApp: +447440014278"),
    }
}

func (m *ManualFiatAdapter) Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    return &PaymentResponse{
        TransactionID: fmt.Sprintf("KONG-%d", time.Now().UnixNano()),
        ProviderTxID:  fmt.Sprintf("MANUAL-%d", time.Now().UnixNano()),
        Status:        "PENDING",
        RedirectURL:   "",
        Metadata: map[string]interface{}{
            "bank_name":      m.bankName,
            "account_name":   m.accountName,
            "account_number": m.accountNumber,
            "instructions":   m.instructions,
            "amount":         req.Amount,
            "currency":       req.Currency,
        },
    }, nil
}

func (m *ManualFiatAdapter) Confirm(ctx context.Context, txID string) (*PaymentStatus, error) {
    return &PaymentStatus{
        Status:    "PENDING",
        SettledAt: time.Now().UTC().Format(time.RFC3339),
    }, nil
}

func (m *ManualFiatAdapter) Cancel(ctx context.Context, txID string) error {
    return nil
}

func (m *ManualFiatAdapter) HandleWebhook(ctx context.Context, payload []byte) (*WebhookEvent, error) {
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

func (m *ManualFiatAdapter) Name() string {
    return "Manual Fiat Transfer"
}

func (m *ManualFiatAdapter) Type() ProviderType {
    return ManualFiat
}

func (m *ManualFiatAdapter) IsAvailable(ctx context.Context) bool {
    return true
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
