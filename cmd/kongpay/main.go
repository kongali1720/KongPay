package main

import (
	"log"

	"github.com/kongali1720/KongPay/internal/config"
	"github.com/kongali1720/KongPay/internal/database"
	"github.com/kongali1720/KongPay/internal/router"
)

func main() {

	cfg := config.Load()

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	r := router.Setup()

	log.Printf(
		"🚀 %s running on :%s (%s)",
		cfg.AppName,
		cfg.Port,
		cfg.Env,
	)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
