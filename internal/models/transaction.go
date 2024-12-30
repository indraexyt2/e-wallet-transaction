package models

import "time"

type Transaction struct {
	ID                int
	UserID            int
	Amount            float64 `gorm:"column:amount;type:decimal(15,2)"`
	TransactionType   string  `gorm:"column:transaction_type;type:enum('TOPUP', 'PURCHASE', 'REFUND')"`
	TransactionStatus string  `gorm:"column:transaction_status;type:enum('PENDING', 'SUCCESS', 'FAILED')"`
	Reference         string  `gorm:"column:reference;type:varchar(255)"`
	Description       string  `gorm:"column:description;type:varchar(255)"`
	AdditionalInfo    string  `gorm:"column:additional_info;type:text"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (t *Transaction) TableName() string {
	return "transactions"
}
