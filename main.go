package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/kongali1720/KongPay/internal/handlers"
    "github.com/kongali1720/KongPay/internal/payment/provider"
    "github.com/kongali1720/KongPay/internal/payment/router"
    "github.com/kongali1720/KongPay/internal/services"
)

func main() {
    log.Println("🚀 KongPay v0.3.0 Starting...")

    // Initialize payment router
    paymentRouter := router.NewPaymentRouter()

    // Register providers
    bankProvider := provider.NewBankAdapter("your-api-key", "https://bank-api.com")
    qrisProvider := provider.NewQRISAdapter("merchant-123", "qris-api-key", "https://qris-api.com")
    cryptoProvider := provider.NewCryptoAdapter("ethereum", "https://rpc.ethereum.org")

    paymentRouter.Register(bankProvider)
    paymentRouter.Register(qrisProvider)
    paymentRouter.Register(cryptoProvider)

    // Initialize services
    txService := services.NewTransactionService()

    // Initialize handlers
    paymentHandler := handlers.NewPaymentHandler(paymentRouter, txService)

    // Setup routes
    http.HandleFunc("/api/v1/payments", paymentHandler.ProcessPayment)
    http.HandleFunc("/api/v1/webhooks/payment", paymentHandler.Webhook)

    // Health check endpoint
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "service":   "KongPay",
            "version":   "0.3.0",
            "status":    "healthy",
            "timestamp": time.Now().UTC().Format(time.RFC3339),
        })
    })

    // Root endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "service": "KongPay",
            "version": "0.3.0",
            "status":  "running",
            "endpoints": []string{
                "GET  /health",
                "GET  /",
                "POST /api/v1/payments",
                "POST /api/v1/webhooks/payment",
            },
        })
    })

    port := ":8080"
    log.Printf("✅ Server running on http://localhost%s", port)
    log.Printf("📊 Health check: http://localhost%s/health", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
