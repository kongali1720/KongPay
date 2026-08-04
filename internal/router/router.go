package router

import (
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"

    "github.com/kongali1720/KongPay/internal/handlers"
    "github.com/kongali1720/KongPay/internal/middleware"
    "github.com/kongali1720/KongPay/internal/payment/provider"
    "github.com/kongali1720/KongPay/internal/payment/router"
    "github.com/kongali1720/KongPay/internal/services"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
    r := gin.Default()

    // Services (with nil DB support)
    txService := services.NewTransactionService(db)

    // Payment Router
    paymentRouter := router.NewPaymentRouter()
    paymentRouter.Register(provider.NewBankAdapter(
        getEnv("BANK_API_KEY", "dummy"),
        getEnv("BANK_BASE_URL", "https://bank-api.com"),
    ))
    paymentRouter.Register(provider.NewQRISAdapter(
        getEnv("QRIS_MERCHANT_ID", "dummy"),
        getEnv("QRIS_API_KEY", "dummy"),
        getEnv("QRIS_BASE_URL", "https://qris-api.com"),
    ))
    paymentRouter.Register(provider.NewCryptoAdapter(
        getEnv("CRYPTO_NETWORK", "ethereum"),
        getEnv("CRYPTO_RPC_URL", "https://rpc.ethereum.org"),
    ))

    // Handlers
    paymentHandler := handlers.NewPaymentHandler(paymentRouter, txService)

    // Routes
    r.GET("/health", handlers.HealthCheck)

    api := r.Group("/api/v1")
    {
        api.POST("/payments", paymentHandler.ProcessPayment)
        api.POST("/webhooks/payment", paymentHandler.Webhook)

        auth := api.Group("/auth")
        {
            auth.POST("/register", handlers.Register)
            auth.POST("/login", handlers.Login)
        }

        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            protected.GET("/wallet", handlers.GetWallet)
            protected.POST("/wallet/topup", handlers.TopUpWallet)
            protected.POST("/wallet/transfer", paymentHandler.Transfer)
        }
    }

    return r
}

func getEnv(key, fallback string) string {
    if value := gin.GetEnv(key); value != "" {
        return value
    }
    return fallback
}
