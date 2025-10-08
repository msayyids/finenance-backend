package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

func (tr *TransactionRepository) AddTransaction(db *sqlx.Tx, transaction *entity.Transaction) error {

	query := `insert into transactions (user_id, category_id,amount,note,created_at,updated_at) values ($1, $2, $3, $4, $5,$6)`

	_, err := db.Exec(query, transaction.UserId, transaction.CategoryId, transaction.Amount, transaction.Note, transaction.CreateAt, transaction.UpdateAt)
	if err != nil {
		return err
	}
	return nil
}
