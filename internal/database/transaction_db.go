package database

import (
	"database/sql"

	"github.com/CelsoTaliatelli/ms-walletconsumer/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (t *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := t.DB.Prepare("INSERT INTO balances (id, account_id_from, account_id_to, amount, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		println("erro database")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transaction.ID, transaction.AccountFromID, transaction.AccountToID, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
