package reconciliation

import (
	"context"

	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
	Repo *repositories.SettlementReconciliationRepository
}

func NewService(
	repo *repositories.SettlementReconciliationRepository,
) *Service {

	return &Service{
		Repo: repo,
	}
}

func (s *Service) Reconcile(
	ctx context.Context,
	settlementID uuid.UUID,
	expected float64,
	actual float64,
) (*models.SettlementReconciliation, error) {

	difference := expected - actual

	status := StatusMatched

	if difference != 0 {
		status = StatusMismatch
	}

	result := &models.SettlementReconciliation{

		ID: uuid.New(),

		SettlementID: settlementID,

		ExpectedAmount: expected,

		ActualAmount: actual,

		Difference: difference,

		Status: status,
	}

	err := s.Repo.Create(
		ctx,
		result,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
