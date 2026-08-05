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
    r := router.SetupRouter(db)

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
