package provider

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
)

type CryptoAdapter struct {
    network string
    rpcURL  string
}

func NewCryptoAdapter(network, rpcURL string) *CryptoAdapter {
    return &CryptoAdapter{
        network: network,
        rpcURL:  rpcURL,
    }
}

func (c *CryptoAdapter) Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    return &PaymentResponse{
        TransactionID: fmt.Sprintf("KONG-%d", time.Now().UnixNano()),
        ProviderTxID:  fmt.Sprintf("CRYPTO-%d", time.Now().UnixNano()),
        Status:        "PENDING",
        QRCode:        "crypto_qr_code_data",
    }, nil
}

func (c *CryptoAdapter) Confirm(ctx context.Context, txID string) (*PaymentStatus, error) {
    return &PaymentStatus{
        Status:    "CONFIRMED",
        SettledAt: time.Now().UTC().Format(time.RFC3339),
    }, nil
}

func (c *CryptoAdapter) Cancel(ctx context.Context, txID string) error {
    return fmt.Errorf("crypto transactions cannot be cancelled")
}

func (c *CryptoAdapter) HandleWebhook(ctx context.Context, payload []byte) (*WebhookEvent, error) {
    var data map[string]interface{}
    if err := json.Unmarshal(payload, &data); err != nil {
        return nil, fmt.Errorf("invalid webhook payload: %w", err)
    }
    return &WebhookEvent{
        TransactionID: getStringValue(data, "transaction_id"),
        ProviderTxID:  getStringValue(data, "tx_hash"),
        Status:        getStringValue(data, "status"),
        Data:          data,
    }, nil
}

func (c *CryptoAdapter) Name() string {
    return "Crypto Provider (" + c.network + ")"
}

func (c *CryptoAdapter) Type() ProviderType {
    return Crypto
}

func (c *CryptoAdapter) IsAvailable(ctx context.Context) bool {
    return true
}

// Helper function
