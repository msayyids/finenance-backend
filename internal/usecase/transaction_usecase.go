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
		Amount:     float64(*request.Amount),
		CategoryId: *request.CategoryId,
		Note:       *request.Note,
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

func (tc *TransactionUseCase) GeTransactionById(transactionId, user_id int) (*entity.Transaction, error) {
	tx, err := tc.DB.Beginx()
	if err != nil {
		tc.Log.Error("failed to create transaction db")
		return nil, errors.New("failed to create transaction")
	}
	defer tx.Rollback()
	transaction, err := tc.TransactionRepo.FindTransactionById(tx, transactionId, user_id)

	if err != nil {
		tc.Log.Error("failed to find transaction", zap.Error(err))
		return nil, errors.New("failed to find transaction")
	}
	err = tx.Commit()
	if err != nil {
		tc.Log.Error("failed to commit", zap.Error(err))
		return nil, errors.New("failed to commit")
	}
	return transaction, nil

}

func (tc *TransactionUseCase) UpdateTransactionById(
	id, user_id int, reqTransaction *model.TransactionRequest,
) (*entity.Transaction, error) {

	tx, err := tc.DB.Beginx()
	if err != nil {
		tc.Log.Error("failed to create transaction db", zap.Error(err))
		return nil, errors.New("failed to create transaction")
	}

	defer tx.Rollback()

	if err := tc.Validator.Struct(reqTransaction); err != nil {
		tc.Log.Warn("invalid user request", zap.Error(err))
		return nil, errors.New("invalid user request")
	}

	transaction, err := tc.TransactionRepo.UpdateTransactionById(tx, id, user_id, reqTransaction)
	if err != nil {
		tc.Log.Error("failed to update transaction", zap.Error(err))
		return nil, errors.New("failed to update transaction")
	}

	err = tx.Commit()
	if err != nil {
		tc.Log.Error("failed to commit", zap.Error(err))
		return nil, errors.New("failed to commit")
	}
	return transaction, nil
}

func (tc *TransactionUseCase) DeleteTransactionById(id, user_id int) error {
	tx, err := tc.DB.Beginx()
	if err != nil {
		tc.Log.Error("failed to create transaction db", zap.Error(err))
		return errors.New("failed to create transaction")
	}

	defer tx.Rollback()

	err = tc.TransactionRepo.DeleteTransactionById(tx, id, user_id)
	if err != nil {
		tc.Log.Error("failed to delete transaction", zap.Error(err))
		return errors.New("failed to delete transaction")
	}
	err = tx.Commit()
	if err != nil {
		tc.Log.Error("failed to commit", zap.Error(err))
		return errors.New("failed to commit")
	}
	return nil

}

func (tc *TransactionUseCase) FindTransactionByIncome(user_id int) (*[]model.TransactionTypeResponse, error) {
	tx, err := tc.DB.Beginx()
	if err != nil {
		tc.Log.Error("failed to create transaction db", zap.Error(err))
		return nil, errors.New("failed to create transaction")
	}

	defer tx.Rollback()

	allTransaction, err := tc.TransactionRepo.FindTransactionByIncome(tx, user_id)
	if err != nil {
		tc.Log.Error("failed to find transaction", zap.Error(err))
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tc.Log.Error("failed to commit", zap.Error(err))
		return nil, errors.New("failed to commit")
	}

	return allTransaction, nil

}
