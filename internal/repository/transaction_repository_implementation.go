package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct{}

type TransactionRepositoryImpl interface {
	AddTransaction(db *sqlx.Tx, transaction *entity.Transaction) error
	FindAllTransactions(db *sqlx.Tx, user_id int) ([]entity.Transaction, error)
}

func NewTransactionRepository() TransactionRepositoryImpl {
	return &TransactionRepository{}
}
