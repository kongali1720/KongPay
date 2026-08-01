package main

import (
	"log"

	"github.com/kongali1720/KongPay/internal/config"
	"github.com/kongali1720/KongPay/internal/router"
)

func main() {

	cfg := config.Load()

	r := router.Setup()

	log.Printf("🚀 %s started on :%s (%s)",
		cfg.AppName,
		cfg.Port,
		cfg.Env,
	)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
