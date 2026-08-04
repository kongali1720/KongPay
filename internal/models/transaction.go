package models

import "time"

type Transaction struct {
    ID             string                 `json:"id"`
    TransactionID  string                 `json:"transaction_id"`
    ProviderTxID   string                 `json:"provider_tx_id"`
    Amount         float64                `json:"amount"`
    Currency       string                 `json:"currency"`
    Method         string                 `json:"method"`
    Status         string                 `json:"status"`
    CustomerID     string                 `json:"customer_id"`
    MerchantID     string                 `json:"merchant_id"`
    RedirectURL    string                 `json:"redirect_url"`
    QRCode         string                 `json:"qr_code"`
    VirtualAccount string                 `json:"virtual_account"`
    Metadata       map[string]interface{} `json:"metadata"`
    CreatedAt      time.Time              `json:"created_at"`
    UpdatedAt      time.Time              `json:"updated_at"`
    SettledAt      *time.Time             `json:"settled_at"`
}
