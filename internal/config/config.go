package config

import "os"

type Config struct {
	AppName string
	Port    string
	Env     string
}

func Load() *Config {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		AppName: "KongPay",
		Port:    port,
		Env:     env,
	}
}
