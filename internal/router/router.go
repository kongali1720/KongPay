package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/handlers"
	"github.com/kongali1720/KongPay/internal/repositories"
	"github.com/kongali1720/KongPay/internal/services"
)

func Setup(db *pgx.Conn) *gin.Engine {

	r := gin.Default()

	repo := repositories.NewWalletRepository(db)
	service := services.NewWalletService(repo)
	handler := handlers.NewWalletHandler(service)

	r.GET("/", handlers.Home)
	r.GET("/health", handlers.Health)

	api := r.Group("/api/v1")
	{
		api.POST("/wallets", handler.CreateWallet)

		api.GET("/wallets", handler.ListWallets)

		api.GET("/wallets/:id", handler.GetWallet)

		api.PUT("/wallets/:id", handler.UpdateWallet)

		api.DELETE("/wallets/:id", handler.DeleteWallet)
	}

	return r
}
