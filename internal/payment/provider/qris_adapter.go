package provider

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
)

type QRISAdapter struct {
    merchantID string
    apiKey     string
}

func NewQRISAdapter(merchantID, apiKey string) *QRISAdapter {
    return &QRISAdapter{
        merchantID: merchantID,
        apiKey:     apiKey,
    }
}

func (q *QRISAdapter) Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    return &PaymentResponse{
        TransactionID: fmt.Sprintf("KONG-%d", time.Now().UnixNano()),
        ProviderTxID:  fmt.Sprintf("QRIS-%d", time.Now().UnixNano()),
        Status:        "PENDING",
        QRCode:        "qris_qr_code_data_here",
    }, nil
}

func (q *QRISAdapter) Confirm(ctx context.Context, txID string) (*PaymentStatus, error) {
    return &PaymentStatus{
        Status: "SUCCESS",
    }, nil
}

func (q *QRISAdapter) Cancel(ctx context.Context, txID string) error {
    return nil
}

func (q *QRISAdapter) HandleWebhook(ctx context.Context, payload []byte) (*WebhookEvent, error) {
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

func (q *QRISAdapter) Name() string {
    return "QRIS Provider"
}

func (q *QRISAdapter) Type() ProviderType {
    return QRIS
}

func (q *QRISAdapter) IsAvailable(ctx context.Context) bool {
    return true
}
