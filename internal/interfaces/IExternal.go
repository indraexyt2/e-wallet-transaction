package interfaces

import (
	"context"
	"e-wallet-transaction/external"
	"e-wallet-transaction/internal/models"
)

type IExternal interface {
	ValidateToken(ctx context.Context, token string) (*models.TokenData, error)
	CreditBalance(ctx context.Context, token string, req external.UpdateBalance) (*external.UpdateBalanceResponse, error)
	DebitBalance(ctx context.Context, token string, req external.UpdateBalance) (*external.UpdateBalanceResponse, error)
	SendNotification(ctx context.Context, recipient string, templateName string, placeholder map[string]string) error
}
