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

	// Wallet
	walletRepo := repositories.NewWalletRepository(db)
	walletService := services.NewWalletService(walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	// User / Auth
	userRepo := repositories.NewRepository(db)
	userService := services.NewService(userRepo, walletRepo)
	authHandler := handlers.NewAuthHandler(userService)

	// Public
	r.GET("/", handlers.Home)
	r.GET("/health", handlers.Health)

	api := r.Group("/api/v1")
	{
		// Authentication
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/login", authHandler.Login)
		api.GET("/auth/profile", authHandler.Profile)

		// Wallet
		api.POST("/wallets", walletHandler.CreateWallet)
		api.GET("/wallets", walletHandler.ListWallets)
		api.GET("/wallets/:id", walletHandler.GetWallet)
		api.PUT("/wallets/:id", walletHandler.UpdateWallet)
		api.DELETE("/wallets/:id", walletHandler.DeleteWallet)
	}

	return r
}
