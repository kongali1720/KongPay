package main

import (
    "log"
    "net/http"

    "github.com/kongali1720/KongPay/internal/handlers"
    "github.com/kongali1720/KongPay/internal/payment/provider"
    "github.com/kongali1720/KongPay/internal/payment/router"
    "github.com/kongali1720/KongPay/internal/services"
)

func main() {
    paymentRouter := router.NewPaymentRouter()
    bankProvider := provider.NewBankAdapter("your-api-key", "https://bank-api.com")
    paymentRouter.Register(bankProvider)

    txService := services.NewTransactionService()
    paymentHandler := handlers.NewPaymentHandler(paymentRouter, txService)

    http.HandleFunc("/api/v1/payments", paymentHandler.ProcessPayment)
    http.HandleFunc("/api/v1/webhooks/payment", paymentHandler.Webhook)

    log.Println("🚀 KongPay Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
