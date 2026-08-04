package router

import (
    "os"

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

    // Services
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

    // Handlers - Buat wrapper Gin
    paymentHandler := handlers.NewPaymentHandler(paymentRouter, txService)

    // Routes
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "service":   "KongPay",
            "version":   "1.0.0-alpha.8.1",
            "status":    "healthy",
            "timestamp": c.Request.Header.Get("Date"),
        })
    })

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "service": "KongPay",
            "version": "1.0.0-alpha.8.1",
            "status":  "running",
            "endpoints": []string{
                "GET  /health",
                "GET  /",
                "POST /api/v1/payments",
                "POST /api/v1/webhooks/payment",
                "GET  /api/v1/settlement/stats",
                "GET  /api/v1/settlement/:transaction_id",
            },
        })
    })

    api := r.Group("/api/v1")
    {
        // Payment routes - wrap with Gin context
        api.POST("/payments", func(c *gin.Context) {
            paymentHandler.ProcessPayment(c.Writer, c.Request)
        })
        api.POST("/webhooks/payment", func(c *gin.Context) {
            paymentHandler.Webhook(c.Writer, c.Request)
        })

        // Settlement routes
        api.GET("/settlement/stats", func(c *gin.Context) {
            stats := txService.GetSettlementStats()
            c.JSON(200, stats)
        })

        api.GET("/settlement/:transaction_id", func(c *gin.Context) {
            txID := c.Param("transaction_id")
            status, err := txService.GetSettlementStatus(txID)
            if err != nil {
                c.JSON(500, gin.H{"error": err.Error()})
                return
            }
            if status == nil {
                c.JSON(404, gin.H{"error": "Settlement not found"})
                return
            }
            c.JSON(200, status)
        })

        // Auth routes
        auth := api.Group("/auth")
        {
            auth.POST("/register", handlers.Register)
            auth.POST("/login", handlers.Login)
        }

        // Protected routes
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            protected.GET("/wallet", handlers.GetWallet)
            protected.POST("/wallet/topup", handlers.TopUpWallet)
            protected.POST("/wallet/transfer", func(c *gin.Context) {
                paymentHandler.Transfer(c)
            })
        }
    }

    return r
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
