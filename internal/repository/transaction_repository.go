package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

func (tr *TransactiomRepository) AddTransaction(db *sqlx.Tx, transaction *entity.Transaction) error {

	query := `
		INSERT INTO transactions (user_id, category_id, amount, note, created_at, updated_at)
		VALUES (:user_id, :category_id, :amount, :note, :created_at, :updated_at)
	`

	_, err := db.NamedExec(query, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TransactiomRepository) FindAllTransactions(db *sqlx.DB, transaction *entity.Transaction) (*[]entity.Transaction, error) {
	query := `SELECT * FROM transactions WHERE user_id = $1 ORDER BY created_at DESC`

	var transactions []entity.Transaction

	err := db.Select(&transactions, query, transaction.UserId)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

//func (tr *TransactiomRepository) FindTransactionById(db *sqlx.DB, transactionId int) (*entity.Transaction, error) {
//
//	query := `SELECT * FROM transactions WHERE id = $1`
//
//	var transaction entity.Transaction
//
//	query := `insert into transactions (user_id, category_id,amount,note,created_at,updated_at) values ($1, $2, $3, $4, $5,$6)`
//
//	_, err := db.Exec(query, transaction.UserId, transaction.CategoryId, transaction.Amount, transaction.Note, transaction.CreateAt, transaction.UpdateAt)
//	if err != nil {
//		return err
//	}
//	return nil
//}
