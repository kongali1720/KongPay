package payment

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
	DB         *pgx.Conn
	WalletRepo *repositories.WalletRepository
	TxRepo     *repositories.TransactionRepository
	LedgerRepo *repositories.LedgerRepository
}

func NewService(
	db *pgx.Conn,
	wallet *repositories.WalletRepository,
	txRepo *repositories.TransactionRepository,
	ledger *repositories.LedgerRepository,
) *Service {

	return &Service{
		DB:         db,
		WalletRepo: wallet,
		TxRepo:     txRepo,
		LedgerRepo: ledger,
	}
}

func (s *Service) Transfer(
	ctx context.Context,
	req TransferRequest,
) (*TransferResponse, error) {

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	sender, err := s.WalletRepo.GetWalletByID(
		ctx,
		req.SenderWalletID,
	)
	if err != nil {
		return nil, err
	}

	receiver, err := s.WalletRepo.GetWalletByID(
		ctx,
		req.ReceiverWalletID,
	)
	if err != nil {
		return nil, err
	}

	if sender.Balance < req.Amount {
		return nil, errors.New("insufficient balance")
	}

	senderBefore := sender.Balance
	receiverBefore := receiver.Balance

	sender.Balance -= req.Amount
	receiver.Balance += req.Amount

	if err := s.WalletRepo.UpdateBalanceTx(
		ctx,
		tx,
		sender.ID,
		sender.Balance,
	); err != nil {
		return nil, err
	}

	if err := s.WalletRepo.UpdateBalanceTx(
		ctx,
		tx,
		receiver.ID,
		receiver.Balance,
	); err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		ID:               uuid.New(),
		ReferenceNo:      GenerateReference(),
		SenderWalletID:   sender.ID,
		ReceiverWalletID: receiver.ID,
		Amount:           req.Amount,
		Currency:         req.Currency,
		Status:           "SUCCESS",
	}

	if err := s.TxRepo.Create(
		ctx,
		tx,
		transaction,
	); err != nil {
		return nil, err
	}

	debit := &models.LedgerEntry{
		ID:            uuid.New(),
		TransactionID: transaction.ID,
		WalletID:      sender.ID,
		EntryType:     "DEBIT",
		Amount:        req.Amount,
		BalanceBefore: senderBefore,
		BalanceAfter:  sender.Balance,
	}

	if err := s.LedgerRepo.Create(
		ctx,
		tx,
		debit,
	); err != nil {
		return nil, err
	}

	credit := &models.LedgerEntry{
		ID:            uuid.New(),
		TransactionID: transaction.ID,
		WalletID:      receiver.ID,
		EntryType:     "CREDIT",
		Amount:        req.Amount,
		BalanceBefore: receiverBefore,
		BalanceAfter:  receiver.Balance,
	}

	if err := s.LedgerRepo.Create(
		ctx,
		tx,
		credit,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &TransferResponse{
		Success:     true,
		ReferenceNo: transaction.ReferenceNo,
		Status:      "SUCCESS",
		Message:     "transfer completed",
	}, nil
}

func GenerateReference() string {
	return "KP-" + uuid.New().String()
}
