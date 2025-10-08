package usecase

import (
	"errors"
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"time"

	"go.uber.org/zap"
)

func (tc *TransactionUseCase) CreateTransaction(user_id int, request *model.TransactionRequest) error {

	tx, err := tc.DB.Beginx()
	if err != nil {
		tc.Log.Error("failed to create transaction db")
		return errors.New("failed to create transaction")

	}

	defer tx.Rollback()

	if err := tc.Validator.Struct(request); err != nil {
		tc.Log.Warn("invalid user request", zap.Error(err))
		return errors.New("invalid user request")
	}

	transaction := entity.Transaction{
		UserId:     user_id,
		Amount:     float64(request.Amount),
		CategoryId: request.CategoryId,
		Note:       request.Note,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}

	err = tc.TransactionRepo.AddTransaction(tx, &transaction)
	if err != nil {
		tc.Log.Error("failed to create new transaction", zap.Error(err))
		return errors.New("failed to create new transaction")
	}

	err = tx.Commit()
	if err != nil {
		tc.Log.Error("failed to commit", zap.Error(err))
		return errors.New("failed to commit")
	}
	return nil
}

func (tc *TransactionUseCase) GetAllTransactions(user_id int) ([]entity.Transaction, error) {
	tx, err := tc.DB.Beginx()
	if err != nil {
		tc.Log.Error("failed to create transaction db")
		return nil, errors.New("failed to create transaction")

	}

	defer tx.Rollback()

	allTransaction, err := tc.TransactionRepo.FindAllTransactions(tx, user_id)
	if err != nil {
		tc.Log.Error("failed to find all transactions", zap.Error(err))
		return nil, errors.New("failed to find all transactions")
	}
	err = tx.Commit()
	if err != nil {
		tc.Log.Error("failed to commit", zap.Error(err))
		return nil, errors.New("failed to commit")
	}

	return allTransaction, nil
}
