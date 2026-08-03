package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/handlers"
	"github.com/kongali1720/KongPay/internal/payment"
	"github.com/kongali1720/KongPay/internal/reconciliation"
	"github.com/kongali1720/KongPay/internal/repositories"
	"github.com/kongali1720/KongPay/internal/services"
	"github.com/kongali1720/KongPay/internal/settlement"
)

func Setup(db *pgx.Conn) *gin.Engine {

	r := gin.Default()

	// Wallet
	walletRepo := repositories.NewWalletRepository(db)
	walletService := services.NewWalletService(walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	// Auth
	userRepo := repositories.NewRepository(db)
	userService := services.NewService(userRepo, walletRepo)
	authHandler := handlers.NewAuthHandler(userService)

	// Payment Engine
	txRepo := repositories.NewTransactionRepository(db)
	ledgerRepo := repositories.NewLedgerRepository()

	paymentService := payment.NewService(
		db,
		walletRepo,
		txRepo,
		ledgerRepo,
	)

	paymentHandler := handlers.NewPaymentHandler(paymentService)

	// Transaction
	transactionHandler := handlers.NewTransactionHandler(txRepo)

	// Ledger
	ledgerHandler := handlers.NewLedgerHandler(
		db,
		ledgerRepo,
	)

	// Settlement
	settlementRepo := repositories.NewSettlementRepository(db)

	settlementEventRepo := repositories.NewSettlementEventRepository(db)

	settlementService := settlement.NewService(
		settlementRepo,
		settlementEventRepo,
	)

	settlementHandler := handlers.NewSettlementHandler(
		settlementService,
	)

	settlementEventHandler := handlers.NewSettlementEventHandler(
		settlementEventRepo,
	)

	// Reconciliation Engine
	reconciliationRepo := repositories.NewSettlementReconciliationRepository(db)

	reconciliationEventRepo := repositories.NewReconciliationEventRepository(db)

	reconciliationService := reconciliation.NewService(
		reconciliationRepo,
		reconciliationEventRepo,
	)

	reconciliationHandler := handlers.NewReconciliationHandler(
		reconciliationService,
	)

	// Public
	r.GET("/", handlers.Home)
	r.GET("/health", handlers.Health)

	api := r.Group("/api/v1")
	{

		// Auth
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/login", authHandler.Login)
		api.GET("/auth/profile", authHandler.Profile)

		// Wallet
		api.POST("/wallets", walletHandler.CreateWallet)
		api.GET("/wallets", walletHandler.ListWallets)
		api.GET("/wallets/:id", walletHandler.GetWallet)
		api.PUT("/wallets/:id", walletHandler.UpdateWallet)
		api.DELETE("/wallets/:id", walletHandler.DeleteWallet)

		// Payment
		api.POST("/transfers", paymentHandler.Transfer)

		// Transaction
		api.GET("/transactions", transactionHandler.List)

		// Ledger
		api.GET("/ledger/:wallet_id", ledgerHandler.ListByWallet)

		// Settlement
		api.POST("/settlements", settlementHandler.Create)
		api.GET("/settlements/:id", settlementHandler.Get)
		api.POST("/settlements/:id/process", settlementHandler.Process)
		api.GET("/settlements/:id/events", settlementEventHandler.List)

		// Reconciliation
		api.POST("/settlements/:id/reconcile", reconciliationHandler.Reconcile)
	}

	return r
}
