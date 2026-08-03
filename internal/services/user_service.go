package services

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/kongali1720/KongPay/internal/auth"
	"github.com/kongali1720/KongPay/internal/models"
	"github.com/kongali1720/KongPay/internal/repositories"
)

type Service struct {
	userRepo   *repositories.Repository
	walletRepo *repositories.WalletRepository
}

func NewService(
	userRepo *repositories.Repository,
	walletRepo *repositories.WalletRepository,
) *Service {
	return &Service{
		userRepo:   userRepo,
		walletRepo: walletRepo,
	}
}

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Phone    string `json:"phone"`
	Currency string `json:"currency"`
}

func (s *Service) Register(req RegisterRequest) (*models.User, error) {

	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: hash,
		Phone:        req.Phone,
		Status:       "ACTIVE",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	currency := req.Currency
	if currency == "" {
		currency = "IDR"
	}

	wallet := &models.Wallet{
		ID:       uuid.New(),
		UserID:   user.ID,
		Balance:  0,
		Currency: currency,
		Status:   "ACTIVE",
	}

	if err := s.walletRepo.CreateWallet(context.Background(), wallet); err != nil {
		return nil, err
	}

	return user, nil
}
