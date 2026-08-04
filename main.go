package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/kongali1720/KongPay/internal/database"
    "github.com/kongali1720/KongPay/internal/router"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  No .env file found, using environment variables")
    }

    // Database config
    dbConfig := database.Config{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     getEnv("DB_PORT", "5432"),
        User:     getEnv("DB_USER", "postgres"),
        Password: getEnv("DB_PASSWORD", "postgres"),
        DBName:   getEnv("DB_NAME", "kongpay"),
        SSLMode:  getEnv("DB_SSL_MODE", "disable"),
    }

    // Connect to database
    db, err := database.NewPostgresDB(dbConfig)
    if err != nil {
        log.Printf("⚠️  Database connection failed: %v", err)
        log.Println("✅ Running without database (development mode)")
    } else {
        defer db.Close()
        log.Println("✅ Database connected successfully!")
    }

    log.Println("🚀 KongPay v1.0.0-alpha.8.1 Starting...")
    log.Println("💰 Settlement Engine: ENABLED")
    log.Println("🔄 Webhook Handler: ENABLED")

    // Setup router
    r := router.SetupRouter(db)

    port := getEnv("PORT", "8080")
    log.Printf("✅ Server running on http://localhost:%s", port)
    log.Printf("📊 Health check: http://localhost:%s/health", port)
    log.Printf("💳 Payment endpoint: http://localhost:%s/api/v1/payments", port)
    log.Printf("🔄 Webhook endpoint: http://localhost:%s/api/v1/webhooks/payment", port)
    log.Printf("💰 Settlement stats: http://localhost:%s/api/v1/settlement/stats", port)

    if err := r.Run(":" + port); err != nil {
        log.Fatalf("❌ Failed to start server: %v", err)
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        return value
    }
    return fallback
}
