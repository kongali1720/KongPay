package settlement

import (
	"context"

	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
	Repo      *repositories.SettlementRepository
	EventRepo *repositories.SettlementEventRepository
}

func NewService(
	repo *repositories.SettlementRepository,
	eventRepo *repositories.SettlementEventRepository,
) *Service {

	return &Service{
		Repo:      repo,
		EventRepo: eventRepo,
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

func (s *Service) Process(
	ctx context.Context,
	id uuid.UUID,
) (*models.Settlement, error) {

	settlement, err := s.Repo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	oldStatus := settlement.Status

	settlement.Status = models.SettlementProcessing

	err = s.EventRepo.Create(
		ctx,
		&models.SettlementEvent{
			ID: uuid.New(),

			SettlementID: settlement.ID,

			EventType: models.SettlementEventProcessing,

			OldStatus: oldStatus,

			NewStatus: settlement.Status,
		},
	)

	if err != nil {
		return nil, err
	}

	// Settlement reconciliation logic
	// transaction matching
	// external payout integration

	oldStatus = settlement.Status

	settlement.Status = models.SettlementCompleted

	err = s.EventRepo.Create(
		ctx,
		&models.SettlementEvent{
			ID: uuid.New(),

			SettlementID: settlement.ID,

			EventType: models.SettlementEventCompleted,

			OldStatus: oldStatus,

			NewStatus: settlement.Status,
		},
	)

	if err != nil {
		return nil, err
	}

	return settlement, nil
}
