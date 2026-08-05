package models

import (
    "time"
)

type Transaction struct {
    ID             string    `json:"id" db:"id"`
    TransactionID  string    `json:"transaction_id" db:"transaction_id"`
    ProviderTxID   string    `json:"provider_tx_id" db:"provider_tx_id"`
    Amount         float64   `json:"amount" db:"amount"`
    Currency       string    `json:"currency" db:"currency"`
    Method         string    `json:"method" db:"method"`
    Status         string    `json:"status" db:"status"`
    CustomerID     string    `json:"customer_id" db:"customer_id"`
    MerchantID     string    `json:"merchant_id" db:"merchant_id"`
    RedirectURL    string    `json:"redirect_url" db:"redirect_url"`
    QRCode         string    `json:"qr_code" db:"qr_code"`
    VirtualAccount string    `json:"virtual_account" db:"virtual_account"`
    Metadata       map[string]interface{} `json:"metadata" db:"metadata"`
    CreatedAt      time.Time `json:"created_at" db:"created_at"`
    UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
    SettledAt      *time.Time `json:"settled_at" db:"settled_at"`
}
