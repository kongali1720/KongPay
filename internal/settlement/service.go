package settlement

import (
	"context"

	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
	Repo *repositories.SettlementRepository
}

func NewService(
	repo *repositories.SettlementRepository,
) *Service {

	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateBatch(
	ctx context.Context,
	currency string,
	totalAmount float64,
	transactionCount int,
) (*models.Settlement, error) {

	settlement := &models.Settlement{
		ID: uuid.New(),

		BatchReference: repositories.GenerateSettlementReference(),

		Currency: currency,

		TotalAmount: totalAmount,

		TransactionCount: transactionCount,

		Status: models.SettlementCreated,
	}

	err := s.Repo.Create(
		ctx,
		settlement,
	)

	if err != nil {
		return nil, err
	}

	return settlement, nil
}

func (s *Service) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Settlement, error) {

	return s.Repo.FindByID(
		ctx,
		id,
	)
}
