package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

type TransactiomRepository struct{}

type TransactiomRepositoryImplementation interface {
	AddTransaction(db *sqlx.Tx, transaction *entity.Transaction) error
	FindAllTransactions(db *sqlx.DB, transaction *entity.Transaction) (*[]entity.Transaction, error)
	FindTransactionById(db *sqlx.DB, transactionId int) (*entity.Transaction, error)
	UpdateTransaction(db *sqlx.Tx, transaction *entity.Transaction) (*entity.Transaction, error)
	DeleteTransactionById(db *sqlx.DB, transactionId, user_id int) error
	FindTransactionsByCategory(db *sqlx.DB, categoryId, user_id int) (*[]entity.Transaction, error)
	FindTransactionsByType(db *sqlx.DB, categoryId, user_id int, types string) (*[]entity.Transaction, error)
}

func NewTransactiomRepository() TransactiomRepositoryImplementation {
	return &TransactiomRepository{}
}
