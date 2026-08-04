package handlers

import (
    "encoding/json"
    "net/http"
    "kongpay/internal/payment/provider"
    "kongpay/internal/payment/router"
    "kongpay/internal/services"
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
    // Get provider from query or header
    providerType := r.URL.Query().Get("provider")
    
    // Read body
    var payload []byte
    r.Body.Read(payload)
    
    // Get provider
    // TODO: Get provider from router
    
    // Process webhook
    // event, err := provider.HandleWebhook(r.Context(), payload)
    // Update transaction status
    
    w.WriteHeader(http.StatusOK)
}
