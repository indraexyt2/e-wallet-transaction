package services

import (
	"context"
	"e-wallet-transaction/constants"
	"e-wallet-transaction/external"
	"e-wallet-transaction/helpers"
	"e-wallet-transaction/internal/interfaces"
	"e-wallet-transaction/internal/models"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type TransactionService struct {
	TransactionRepo interfaces.ITransactionRepository
	External        interfaces.IExternal
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *models.Transaction) (*models.CreateTransactionResponse, error) {
	var (
		resp = &models.CreateTransactionResponse{}
	)

	req.TransactionStatus = constants.TransactionStatusPending
	req.Reference = helpers.GenerateReference()

	jsonAdditionalIndo := map[string]interface{}{}
	if req.AdditionalInfo != "" {
		err := json.Unmarshal([]byte(req.AdditionalInfo), &jsonAdditionalIndo)
		if err != nil {
			return resp, errors.Wrap(err, "additional info type is invalid")
		}
	}

	jsonAdditionalInfo := map[string]interface{}{}
	if req.AdditionalInfo != "" {
		err := json.Unmarshal([]byte(req.AdditionalInfo), &jsonAdditionalInfo)
		if err != nil {
			return resp, errors.Wrap(err, "additional info type is invalid")
		}
	}

	err := s.TransactionRepo.CreateTransaction(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "failed to create transaction")
	}

	resp.Reference = req.Reference
	resp.TransactionStatus = req.TransactionStatus
	return resp, nil
}

func (s *TransactionService) UpdateStatusTransaction(ctx context.Context, tokenData *models.TokenData, req *models.UpdateStatusTransaction) error {
	trx, err := s.TransactionRepo.GetTransactionByReference(ctx, req.Reference, false)
	if err != nil {
		return errors.Wrap(err, "failed to get transaction by reference")
	}

	statusValid := false
	mapStatusFlow := constants.MapTransactionStatusFlow[trx.TransactionStatus]
	for i := range mapStatusFlow {
		if mapStatusFlow[i] == req.TransactionStatus {
			statusValid = true
		}
	}

	if !statusValid {
		return fmt.Errorf("invalid transaction status: %s", req.TransactionStatus)
	}

	reqUpdateBalance := external.UpdateBalance{
		Reference: req.Reference,
		Amount:    trx.Amount,
	}

	if req.TransactionStatus == constants.TransactionStatusReversed {
		reqUpdateBalance.Reference = "REVERSED-" + req.Reference
		now := time.Now()
		expiredReversalTime := trx.CreatedAt.Add(constants.MaximumReversalDuration)
		if now.After(expiredReversalTime) {
			return errors.New("failed to reverse transaction due to maximum reversal duration")
		}
	}

	var (
		errUpdateBalance error
	)
	switch trx.TransactionType {
	case constants.TransactionTypeTopup:
		if req.TransactionStatus == constants.TransactionStatusSuccess {
			_, errUpdateBalance = s.External.CreditBalance(ctx, tokenData.Token, reqUpdateBalance)
		} else if req.TransactionStatus == constants.TransactionStatusReversed {
			_, errUpdateBalance = s.External.DebitBalance(ctx, tokenData.Token, reqUpdateBalance)
		}
	case constants.TransactionTypePurchase:
		if req.TransactionStatus == constants.TransactionStatusSuccess {
			_, errUpdateBalance = s.External.DebitBalance(ctx, tokenData.Token, reqUpdateBalance)
		} else if req.TransactionStatus == constants.TransactionStatusReversed {
			_, errUpdateBalance = s.External.CreditBalance(ctx, tokenData.Token, reqUpdateBalance)
		}
	}

	if errUpdateBalance != nil {
		return errors.Wrap(errUpdateBalance, "failed to update balance")
	}

	var (
		currentAdditionalInfo = map[string]interface{}{}
		newAdditionalInfo     = map[string]interface{}{}
	)

	if trx.AdditionalInfo != "" {
		err = json.Unmarshal([]byte(trx.AdditionalInfo), &currentAdditionalInfo)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal current additional info")
		}
	}

	if req.AdditionalInfo != "" {
		err = json.Unmarshal([]byte(req.AdditionalInfo), &newAdditionalInfo)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal new additional info")
		}
	}

	for key, value := range newAdditionalInfo {
		currentAdditionalInfo[key] = value
	}

	byteAdditionalInfo, err := json.Marshal(currentAdditionalInfo)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal additional info")
	}

	err = s.TransactionRepo.UpdateStatusTransaction(ctx, req.Reference, req.TransactionStatus, string(byteAdditionalInfo))
	if err != nil {
		return errors.Wrap(err, "failed to update status transaction")
	}

	return nil
}

func (s *TransactionService) GetTransaction(ctx context.Context, userID int) ([]models.Transaction, error) {
	return s.TransactionRepo.GetTransaction(ctx, userID)
}

func (s *TransactionService) GetTransactionDetail(ctx context.Context, reference string) (*models.Transaction, error) {
	return s.TransactionRepo.GetTransactionByReference(ctx, reference, true)
}
