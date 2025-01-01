package interfaces

import (
	"context"
	"e-wallet-transaction/internal/models"
	"github.com/gin-gonic/gin"
)

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, trx *models.Transaction) error
	GetTransactionByReference(ctx context.Context, reference string, includeRefund bool) (*models.Transaction, error)
	UpdateStatusTransaction(ctx context.Context, reference string, status string, additionalInfo string) error
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, req *models.Transaction) (*models.CreateTransactionResponse, error)
	UpdateStatusTransaction(ctx context.Context, tokenData *models.TokenData, req *models.UpdateStatusTransaction) error
}

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	UpdateTransactionStatus(c *gin.Context)
}
