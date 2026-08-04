package services

import (
    "context"
    "log"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/kongali1720/KongPay/internal/models"
    "github.com/kongali1720/KongPay/internal/repositories"
)

type TransactionService struct {
    repo              *repositories.TransactionRepository
    settlementService *SettlementService
}

func NewTransactionService(db *pgxpool.Pool) *TransactionService {
    var txRepo *repositories.TransactionRepository
    if db != nil {
        txRepo = repositories.NewTransactionRepository(db)
        log.Println("✅ Transaction repository initialized with DB")
    } else {
        log.Println("⚠️ Database is nil!")
    }

    return &TransactionService{
        repo:              txRepo,
        settlementService: NewSettlementService(db),
    }
}

func (s *TransactionService) SaveTransaction(ctx context.Context, tx *models.Transaction) error {
    if s.repo == nil {
        log.Printf("❌ REPO NIL: Transaction %s NOT saved!", tx.TransactionID)
        return nil
    }
    log.Printf("💾 Saving transaction: %s", tx.TransactionID)
    err := s.repo.Create(ctx, tx)
    if err != nil {
        log.Printf("❌ Error saving transaction: %v", err)
        return err
    }
    log.Printf("✅ Transaction saved: %s", tx.TransactionID)
    return nil
}

func (s *TransactionService) UpdateTransactionStatus(ctx context.Context, transactionID string, status string) error {
    if s.repo == nil {
        log.Printf("❌ REPO NIL: Status for %s NOT updated!", transactionID)
        return nil
    }
    log.Printf("📝 Updating transaction %s to status: %s", transactionID, status)
    return s.repo.UpdateStatus(ctx, transactionID, status)
}

func (s *TransactionService) GetTransaction(ctx context.Context, transactionID string) (*models.Transaction, error) {
    if s.repo == nil {
        log.Printf("❌ REPO NIL: Cannot get transaction %s", transactionID)
        return nil, nil
    }
    log.Printf("🔍 Getting transaction: %s", transactionID)
    tx, err := s.repo.GetByID(ctx, transactionID)
    if err != nil {
        log.Printf("❌ Error getting transaction: %v", err)
        return nil, err
    }
    if tx == nil {
        log.Printf("❌ Transaction NOT FOUND in DB: %s", transactionID)
        return nil, nil
    }
    log.Printf("✅ Transaction found: %s (Status: %s)", tx.TransactionID, tx.Status)
    return tx, nil
}

func (s *TransactionService) TriggerSettlement(ctx context.Context, transactionID string) {
    log.Printf("💰💰💰 TRIGGERING SETTLEMENT FOR: %s", transactionID)

    tx, err := s.GetTransaction(ctx, transactionID)
    if err != nil {
        log.Printf("❌ Error getting transaction: %v", err)
        return
    }
    if tx == nil {
        log.Printf("❌❌❌ Transaction NOT FOUND: %s", transactionID)
        return
    }

    s.TriggerSettlementDirect(ctx, transactionID, tx.Amount, tx.Currency, tx.CustomerID, tx.MerchantID)
}

func (s *TransactionService) TriggerSettlementDirect(ctx context.Context, transactionID string, amount float64, currency string, customerID string, merchantID string) error {
    log.Printf("💰 Adding settlement task for %s (Amount: %.2f %s)", transactionID, amount, currency)
    
    if err := s.settlementService.AddTask(ctx, transactionID, amount, currency, customerID, merchantID); err != nil {
        log.Printf("❌ Failed to add settlement task: %v", err)
        return err
    }

    s.UpdateTransactionStatus(ctx, transactionID, "SETTLING")
    log.Printf("✅✅✅ Settlement triggered for: %s", transactionID)
    return nil
}

func (s *TransactionService) GetSettlementStatus(transactionID string) (*SettlementTask, error) {
    return s.settlementService.GetStatus(transactionID)
}

func (s *TransactionService) GetSettlementStats() map[string]interface{} {
    return s.settlementService.GetStats()
}
