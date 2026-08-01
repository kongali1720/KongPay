package router

import (
	"github.com/gin-gonic/gin"

	"github.com/kongali1720/KongPay/internal/handlers"
)

func Setup() *gin.Engine {

	r := gin.Default()

	r.GET("/", handlers.Home)

	r.GET("/health", handlers.Health)

	api := r.Group("/api/v1")
	{
		api.POST("/wallets", handlers.CreateWallet)
	}

	return r
}
