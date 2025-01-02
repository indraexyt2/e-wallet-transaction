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
	GetTransaction(ctx context.Context, userID int) ([]models.Transaction, error)
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, req *models.Transaction) (*models.CreateTransactionResponse, error)
	UpdateStatusTransaction(ctx context.Context, tokenData *models.TokenData, req *models.UpdateStatusTransaction) error
	GetTransaction(ctx context.Context, userID int) ([]models.Transaction, error)
	GetTransactionDetail(ctx context.Context, reference string) (*models.Transaction, error)
	RefundTransaction(ctx context.Context, tokenData *models.TokenData, req *models.RefundTransaction) (*models.CreateTransactionResponse, error)
}

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	UpdateTransactionStatus(c *gin.Context)
	GetTransaction(c *gin.Context)
	GetTransactionDetail(c *gin.Context)
	RefundTransaction(c *gin.Context)
}
