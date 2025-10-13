package usecase

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TransactionUseCase struct {
	DB              *sqlx.DB
	Log             *zap.Logger
	Validator       *validator.Validate
	TransactionRepo repository.TransactionRepositoryImpl
}

type TransactionUseCaseImpl interface {
	CreateTransaction(user_id int, request *model.TransactionRequest) error
	GetAllTransactions(user_id int) ([]entity.Transaction, error)
	GeTransactionById(transactionId, user_id int) (*entity.Transaction, error)
	UpdateTransactionById(id, user_id int, reqTransaction *model.TransactionRequest) (*entity.Transaction, error)
	DeleteTransactionById(id, user_id int) error
	FindTransactionByIncome(user_id int) (*[]model.TransactionTypeResponse, error)
	FindTransactionByExpense(user_id int) (*[]model.TransactionTypeResponse, error)
}

func NewTransactionUseCase(
	db *sqlx.DB, log *zap.Logger, Validator *validator.Validate, tr repository.TransactionRepositoryImpl,
) TransactionUseCaseImpl {
	return &TransactionUseCase{
		DB:              db,
		Log:             log,
		Validator:       Validator,
		TransactionRepo: tr,
	}
}
