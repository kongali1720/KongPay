package provider

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

type BankAdapter struct {
    apiKey     string
    baseURL    string
    httpClient *http.Client
}

func NewBankAdapter(apiKey, baseURL string) *BankAdapter {
    return &BankAdapter{
        apiKey:     apiKey,
        baseURL:    baseURL,
        httpClient: &http.Client{Timeout: 30 * time.Second},
    }
}

type BankPaymentRequest struct {
    Amount     float64 `json:"amount"`
    Currency   string  `json:"currency"`
    Reference  string  `json:"reference"`
    CustomerID string  `json:"customer_id"`
    MerchantID string  `json:"merchant_id"`
}

type BankPaymentResponse struct {
    TransactionID string `json:"transaction_id"`
    RedirectURL   string `json:"redirect_url"`
    Status        string `json:"status"`
    Error         string `json:"error,omitempty"`
}

func (b *BankAdapter) Process(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    // Build request to bank
    bankReq := BankPaymentRequest{
        Amount:     req.Amount,
        Currency:   req.Currency,
        Reference:  fmt.Sprintf("KONG-%d", time.Now().UnixNano()),
        CustomerID: req.CustomerID,
        MerchantID: req.MerchantID,
    }

    jsonData, err := json.Marshal(bankReq)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal request: %w", err)
   }

    // Create HTTP request
    httpReq, err := http.NewRequestWithContext(ctx, "POST", b.baseURL+"/api/v1/payments", bytes.NewReader(jsonData))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
   }

    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("X-API-Key", b.apiKey)

    // Send request
    resp, err := b.httpClient.Do(httpReq)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response: %w", err)
    }

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
        return nil, fmt.Errorf("bank API error: %s - %s", resp.Status, string(body))
    }

    var bankResp BankPaymentResponse
    if err := json.Unmarshal(body, &bankResp); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    return &PaymentResponse{
        TransactionID: fmt.Sprintf("KONG-%d", time.Now().UnixNano()),
        ProviderTxID:  bankResp.TransactionID,
        Status:        bankResp.Status,
        RedirectURL:   bankResp.RedirectURL,
    }, nil
}

// ... rest of methods remain the same ...
