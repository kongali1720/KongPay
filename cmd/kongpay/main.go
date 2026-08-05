package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/kongali1720/KongPay/internal/database"
    "github.com/kongali1720/KongPay/internal/handlers"
    "github.com/kongali1720/KongPay/internal/middleware"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️ No .env file found, using environment variables")
    }

    // Database connection
    db, err := database.NewConnection()
    if err != nil {
        log.Printf("⚠️ Database connection failed: %v", err)
        log.Println("✅ Running without database (development mode)")
    } else {
        defer db.Close()
        log.Println("✅ Database connected successfully!")
    }

    // Setup router
    r := gin.Default()

    // Health check
    r.GET("/health", handlers.HealthCheck)

    // API v1
    api := r.Group("/api/v1")
    {
        api.POST("/payments", handlers.ProcessPayment)
        api.POST("/webhooks/payment", handlers.Webhook)
        api.GET("/settlement/stats", handlers.SettlementStats)
        api.GET("/settlement/:transaction_id", handlers.SettlementStatus)

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
            protected.POST("/wallet/transfer", handlers.Transfer)
        }
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("✅ Server running on http://localhost:%s", port)
    log.Printf("📊 Health check: http://localhost:%s/health", port)

    if err := r.Run(":" + port); err != nil {
        log.Fatalf("❌ Failed to start server: %v", err)
    }
}
