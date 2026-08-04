package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

func main() {
    log.Println("🚀 KongPay v0.3.0")

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "healthy",
            "service": "KongPay",
            "version": "0.3.0",
            "timestamp": time.Now().UTC().Format(time.RFC3339),
        })
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "service": "KongPay",
            "version": "0.3.0",
            "status": "running",
        })
    })

    log.Println("✅ Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
