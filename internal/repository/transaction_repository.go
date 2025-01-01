package repository

import (
	"context"
	"e-wallet-transaction/constants"
	"e-wallet-transaction/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, trx *models.Transaction) error {
	return r.DB.Create(trx).Error
}

func (r *TransactionRepository) GetTransactionByReference(ctx context.Context, reference string, includeRefund bool) (*models.Transaction, error) {
	var (
		resp models.Transaction
	)
	sql := r.DB.Where("reference = ?", reference)
	if includeRefund {
		sql.Where("transaction_status != ?", constants.TransactionTypeRefund)
	}
	err := sql.Last(&resp).Error
	if err != nil {
		return &resp, errors.Wrap(err, "failed to get transaction by reference")
	}
	return &resp, nil
}

func (r *TransactionRepository) UpdateStatusTransaction(ctx context.Context, reference string, status string, additionalInfo string) error {
	return r.DB.Model(&models.Transaction{}).
		Where("reference = ?", reference).
		Updates(map[string]interface{}{"transaction_status": status, "additional_info": additionalInfo}).Error
}

func (r *TransactionRepository) GetTransaction(ctx context.Context, userID int) ([]models.Transaction, error) {
	var (
		resp []models.Transaction
	)
	err := r.DB.Where("user_id = ?", userID).Find(&resp).Order("id DESC").Error
	return resp, err
}
