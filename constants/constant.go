package constants

import "time"

const (
	TransactionStatusPending  = "PENDING"
	TransactionStatusSuccess  = "SUCCESS"
	TransactionStatusFailed   = "FAILED"
	TransactionStatusReversed = "REVERSED"
)

const (
	TransactionTypeTopup    = "TOPUP"
	TransactionTypePurchase = "PURCHASE"
	TransactionTypeRefund   = "REFUND"
)

var MapTransaction = map[string]bool{
	TransactionTypeTopup:    true,
	TransactionTypePurchase: true,
	TransactionTypeRefund:   true,
}

var MapTransactionStatusFlow = map[string][]string{
	TransactionStatusPending: {TransactionStatusSuccess, TransactionStatusFailed},
	TransactionStatusSuccess: {TransactionStatusReversed},
	TransactionStatusFailed:  {TransactionStatusSuccess},
}

const (
	MaximumReversalDuration = time.Minute * 2
)

const (
	SuccessMessage = "success"
)
