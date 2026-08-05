package router

import (
    "github.com/gin-gonic/gin"
    "github.com/kongali1720/KongPay/internal/database"
    "github.com/kongali1720/KongPay/internal/handlers"
    "github.com/kongali1720/KongPay/internal/middleware"
)

func SetupRouter(db *database.DB) *gin.Engine {
    r := gin.Default()

    // Health check
    r.GET("/health", handlers.HealthCheck)

    // API v1
    api := r.Group("/api/v1")
    {
        // Payment routes
        api.POST("/payments", handlers.ProcessPayment)
        api.POST("/webhooks/payment", handlers.Webhook)

        // Settlement routes
        api.GET("/settlement/stats", handlers.SettlementStats)
        api.GET("/settlement/:transaction_id", handlers.SettlementStatus)

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
            protected.POST("/wallet/transfer", handlers.Transfer)
        }
    }

    return r
}
