package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/kongali1720/KongPay/internal/payment/provider"
    "github.com/kongali1720/KongPay/internal/payment/router"
    "github.com/kongali1720/KongPay/internal/services"
)

type PaymentHandler struct {
    paymentRouter *router.PaymentRouter
    txService     *services.TransactionService
}

func NewPaymentHandler(paymentRouter *router.PaymentRouter, txService *services.TransactionService) *PaymentHandler {
    return &PaymentHandler{
        paymentRouter: paymentRouter,
        txService:     txService,
    }
}

// ProcessPayment handles payment requests
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

    paymentReq := &provider.PaymentRequest{
        Amount:     req.Amount,
        Currency:   req.Currency,
        Method:     req.Method,
        CustomerID: req.CustomerID,
        MerchantID: req.MerchantID,
        Metadata:   make(map[string]interface{}),
    }

    resp, err := h.paymentRouter.Route(r.Context(), paymentReq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// Webhook handles payment webhooks
func (h *PaymentHandler) Webhook(w http.ResponseWriter, r *http.Request) {
    var payload map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Invalid payload", http.StatusBadRequest)
        return
    }

    transactionID, _ := payload["transaction_id"].(string)
    status, _ := payload["status"].(string)

    log.Printf("📨 Webhook received: TX=%s, Status=%s", transactionID, status)

    if err := h.txService.UpdateTransactionStatus(r.Context(), transactionID, status); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if status == "SUCCESS" {
        go h.txService.TriggerSettlement(r.Context(), transactionID)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status":  "webhook_processed",
        "message": "Webhook processed successfully",
    })
}

// Transfer handles transfer requests (Gin handler)
func (h *PaymentHandler) Transfer(c *gin.Context) {
    var req struct {
        FromWalletID string  `json:"from_wallet_id"`
        ToWalletID   string  `json:"to_wallet_id"`
        Amount       float64 `json:"amount"`
        Currency     string  `json:"currency"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // TODO: Implement transfer logic
    c.JSON(http.StatusOK, gin.H{
        "status":      "success",
        "message":     "Transfer processed",
        "from":        req.FromWalletID,
        "to":          req.ToWalletID,
        "amount":      req.Amount,
        "currency":    req.Currency,
        "transfer_id": "TRF-" + req.FromWalletID[:8],
    })
}
