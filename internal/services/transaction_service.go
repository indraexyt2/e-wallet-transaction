package services

import (
	"context"
	"e-wallet-transaction/constants"
	"e-wallet-transaction/helpers"
	"e-wallet-transaction/internal/interfaces"
	"e-wallet-transaction/internal/models"
	"github.com/pkg/errors"
)

type TransactionService struct {
	TransactionRepo interfaces.ITransactionRepository
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *models.Transaction) (*models.CreateTransactionResponse, error) {
	var (
		resp = &models.CreateTransactionResponse{}
	)

	req.TransactionStatus = constants.TransactionStatusPending
	req.Reference = helpers.GenerateReference()
	err := s.TransactionRepo.CreateTransaction(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "failed to create transaction")
	}

	resp.Reference = req.Reference
	resp.TransactionStatus = req.TransactionStatus
	return resp, nil
}
