package main

import (
    "log"
    "net/http"
    
    "kongpay/internal/handlers"
    "kongpay/internal/payment/provider"
    "kongpay/internal/payment/router"
    "kongpay/internal/services"
)

func main() {
    // Initialize payment router
    paymentRouter := router.NewPaymentRouter()
    
    // Register providers
    bankProvider := provider.NewBankAdapter("your-api-key", "https://bank-api.com")
    paymentRouter.Register(bankProvider)
    
    // Initialize services
    txService := services.NewTransactionService()
    
    // Initialize handlers
    paymentHandler := handlers.NewPaymentHandler(paymentRouter, txService)
    
    // Setup routes
    http.HandleFunc("/api/v1/payments", paymentHandler.ProcessPayment)
    http.HandleFunc("/api/v1/webhooks/payment", paymentHandler.Webhook)
    
    // Start server
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
