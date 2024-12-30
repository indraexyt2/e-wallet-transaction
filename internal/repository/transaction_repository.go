package repository

import (
	"context"
	"e-wallet-transaction/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, trx *models.Transaction) error {
	return r.DB.Create(trx).Error
}
