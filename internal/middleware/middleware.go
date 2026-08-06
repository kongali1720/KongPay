package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware DIHAPUS dari sini.
// Auth wajib pakai auth.JWTMiddleware() dari package internal/auth,
// yang benar-benar memvalidasi signature JWT.
// Jangan buat middleware auth kedua lagi di file ini atau file lain -
// satu sumber kebenaran untuk auth = internal/auth.

// CORSMiddleware membatasi origin yang boleh akses API.
// Ganti allowedOrigins sesuai domain frontend/production Anda.
func CORSMiddleware() gin.HandlerFunc {
    allowedOrigins := map[string]bool{
        "https://kongpay.id":           true,
        "http://localhost:3000":        true, // hapus saat production
    }

    return func(c *gin.Context) {
        origin := c.GetHeader("Origin")
        if allowedOrigins[origin] {
            c.Header("Access-Control-Allow-Origin", origin)
        }
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Idempotency-Key")
        c.Header("Vary", "Origin")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

// RateLimitMiddleware masih placeholder - JANGAN pasang ini di endpoint
// payment/wallet sampai ada implementasi Redis-based rate limiter.
// Lihat patch terpisah untuk versi yang benar-benar membatasi request.
func RateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // TODO: implement - lihat internal/middleware/ratelimit.go (patch berikutnya)
        c.Next()
    }
}
