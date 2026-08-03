package main

import (
	"log"

	"github.com/kongali1720/KongPay/internal/config"
	"github.com/kongali1720/KongPay/internal/database"
	"github.com/kongali1720/KongPay/internal/router"
)

func main() {

	cfg := config.Load()

	log.Printf("Starting %s...", cfg.AppName)

	if err := database.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()

	r := router.Setup(database.DB)

	log.Printf("Listening on :%s", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
