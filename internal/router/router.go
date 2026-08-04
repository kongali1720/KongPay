package router
import ("github.com/gin-gonic/gin"; "github.com/kongali1720/KongPay/internal/handlers"; "github.com/kongali1720/KongPay/internal/middleware"; "github.com/kongali1720/KongPay/internal/payment/provider"; "github.com/kongali1720/KongPay/internal/payment/router"; "github.com/kongali1720/KongPay/internal/services")
func SetupRouter() *gin.Engine {
    r := gin.Default()
    paymentRouter := router.NewPaymentRouter()
    paymentRouter.Register(provider.NewBankAdapter("your-api-key", "https://bank-api.com"))
    paymentRouter.Register(provider.NewQRISAdapter("merchant-123", "qris-api-key", "https://qris-api.com"))
    paymentRouter.Register(provider.NewCryptoAdapter("ethereum", "https://rpc.ethereum.org"))
    txService := services.NewTransactionService()
    paymentHandler := handlers.NewPaymentHandler(paymentRouter, txService)
    r.GET("/health", handlers.HealthCheck)
    api := r.Group("/api/v1")
    {
        api.POST("/payments", paymentHandler.ProcessPayment)
        api.POST("/webhooks/payment", paymentHandler.Webhook)
        auth := api.Group("/auth")
        {
            auth.POST("/register", handlers.Register)
            auth.POST("/login", handlers.Login)
        }
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            protected.GET("/wallet", handlers.GetWallet)
            protected.POST("/wallet/topup", handlers.TopUpWallet)
        }
    }
    return r
}
