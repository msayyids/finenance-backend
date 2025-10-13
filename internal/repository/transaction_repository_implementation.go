package repository

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct{}

type TransactionRepositoryImpl interface {
	AddTransaction(db *sqlx.Tx, transaction *entity.Transaction) error
	FindAllTransactions(db *sqlx.Tx, user_id int) ([]entity.Transaction, error)
	FindTransactionById(db *sqlx.Tx, transactionId, user_id int) (*entity.Transaction, error)
	UpdateTransactionById(db *sqlx.Tx, id, user_id int, reqTransaction *model.TransactionRequest) (
		*entity.Transaction, error,
	)
	DeleteTransactionById(db *sqlx.Tx, id, user_id int) error
	FindTransactionByIncome(db *sqlx.Tx, userId int) (*[]model.TransactionTypeResponse, error)
	FindTransactionByExpense(db *sqlx.Tx, userId int) (*[]model.TransactionTypeResponse, error)
}

func NewTransactionRepository() TransactionRepositoryImpl {
	return &TransactionRepository{}
}
