package handlers

import (
    "encoding/json"
    "net/http"
    
    "github.com/kongali1720/KongPay/internal/payment/provider"
    "github.com/kongali1720/KongPay/internal/payment/router"
    "github.com/kongali1720/KongPay/internal/services"
)

type PaymentHandler struct {
    paymentRouter *router.PaymentRouter
    txService     *services.TransactionService
}

func NewPaymentHandler(router *router.PaymentRouter, txService *services.TransactionService) *PaymentHandler {
    return &PaymentHandler{
        paymentRouter: router,
        txService:     txService,
    }
}

func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Amount     float64 `json:"amount"`
        Currency   string  `json:"currency"`
        Method     string  `json:"method"`
        CustomerID string  `json:"customer_id"`
        MerchantID string  `json:"merchant_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create payment request
    paymentReq := &provider.PaymentRequest{
        Amount:     req.Amount,
        Currency:   req.Currency,
        Method:     req.Method,
        CustomerID: req.CustomerID,
        MerchantID: req.MerchantID,
        Metadata:   make(map[string]interface{}),
    }

    // Route to provider
    resp, err := h.paymentRouter.Route(r.Context(), paymentReq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *PaymentHandler) Webhook(w http.ResponseWriter, r *http.Request) {
    // Read body
    var payload []byte
    if _, err := r.Body.Read(payload); err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }

    // TODO: Process webhook based on provider
    // Get provider from query: r.URL.Query().Get("provider")
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "webhook_received",
    })
}
