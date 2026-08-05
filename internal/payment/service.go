package payment

import (
    "context"
    "time"

    "github.com/google/uuid"
    "github.com/kongali1720/KongPay/internal/models"
    "github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
    TxRepo *repositories.TransactionRepository
}

func NewService(txRepo *repositories.TransactionRepository) *Service {
    return &Service{
        TxRepo: txRepo,
    }
}

func (s *Service) CreateTransaction(ctx context.Context, amount float64, currency, method, customerID, merchantID string) (*models.Transaction, error) {
    transaction := &models.Transaction{
        ID:           uuid.New().String(),
        TransactionID: "KONG-" + uuid.New().String()[:8],
        Amount:       amount,
        Currency:     currency,
        Method:       method,
        Status:       "PENDING",
        CustomerID:   customerID,
        MerchantID:   merchantID,
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }

    if err := s.TxRepo.Create(ctx, transaction); err != nil {
        return nil, err
    }

    return transaction, nil
}

func (s *Service) GetTransaction(ctx context.Context, transactionID string) (*models.Transaction, error) {
    return s.TxRepo.GetByID(ctx, transactionID)
}

func (s *Service) UpdateTransactionStatus(ctx context.Context, transactionID, status string) error {
    return s.TxRepo.UpdateStatus(ctx, transactionID, status)
}
