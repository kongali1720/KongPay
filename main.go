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
        Password: getEnv("DB_PASSWORD", "password"),
        DBName:   getEnv("DB_NAME", "kongpay"),
        SSLMode:  getEnv("DB_SSL_MODE", "disable"),
    }

    // Connect to database
    db, err := database.NewPostgresDB(dbConfig)
    if err != nil {
        log.Printf("⚠️  Database connection failed: %v", err)
        log.Println("✅ Running without database (development mode)")
        // Continue without DB for now
    } else {
        defer db.Close()
        log.Println("✅ Database connected")
    }

    log.Println("🚀 KongPay v0.4.0 Starting...")

    // Setup router
    r := router.SetupRouter(db)

    port := getEnv("PORT", "8080")
    log.Printf("✅ Server running on http://localhost:%s", port)
    log.Printf("📊 Health check: http://localhost:%s/health", port)

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
