package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/kongali1720/KongPay/internal/database"
    "github.com/kongali1720/KongPay/internal/router"
)

func main() {
    // Load config
    loadConfig()

    // Check if database is disabled
    dbDisabled := getEnv("DB_DISABLE", "false") == "true"

    var dbConn interface{} // Gunakan interface{} agar fleksibel
    if !dbDisabled {
        dbConfig := database.Config{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnv("DB_PORT", "5432"),
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            DBName:   getEnv("DB_NAME", "kongpay"),
            SSLMode:  getEnv("DB_SSL_MODE", "disable"),
        }

        db, err := database.NewPostgresDB(dbConfig)
        if err != nil {
            log.Printf("⚠️  Database connection failed: %v", err)
            log.Println("⚠️  Running without database")
        } else {
            defer db.Close()
            log.Println("✅ Database connected successfully!")
            dbConn = db
        }
    } else {
        log.Println("⚠️  Database disabled by DB_DISABLE=true")
    }

    log.Println("🚀 KongPay v1.0.0-alpha.8.1 Starting...")
    log.Println("💰 Settlement Engine: ENABLED")
    log.Println("🔄 Webhook Handler: ENABLED")

    // Setup router (handle nil db)
    r := router.SetupRouter(dbConn)

    port := getEnv("PORT", "8080")
    log.Printf("✅ Server running on http://localhost:%s", port)
    log.Printf("📊 Health check: http://localhost:%s/health", port)

    if err := r.Run(":" + port); err != nil {
        log.Fatalf("❌ Failed to start server: %v", err)
    }
}

func loadConfig() {
    configDir := os.Getenv("KONGPAY_CONFIG")
    if configDir == "" {
        configDir = os.ExpandEnv("$HOME/KongPay-Config")
    }

    envFiles := []string{
        configDir + "/.env",
        configDir + "/.env.fiat",
        configDir + "/.env.crypto",
    }

    for _, file := range envFiles {
        if _, err := os.Stat(file); err == nil {
            if err := godotenv.Load(file); err == nil {
                log.Printf("✅ Loaded config: %s", file)
            }
        }
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        return value
    }
    return fallback
}
