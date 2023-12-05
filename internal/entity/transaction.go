package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            string
	AccountFromID string `gorm:"foreignKey:AccountFromID"`
	AccountToID   string `gorm:"foreignKey:AccountToID"`
	Amount        float64
	CreatedAt     time.Time
}

func NewTransaction(accountFrom string, accountTo string, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:            uuid.New().String(),
		AccountFromID: accountFrom,
		AccountToID:   accountTo,
		Amount:        amount,
		CreatedAt:     time.Now(),
	}

	return transaction, nil
}
