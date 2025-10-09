package repository

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"

	"github.com/jmoiron/sqlx"
)

func (tr *TransactionRepository) AddTransaction(db *sqlx.Tx, transaction *entity.Transaction) error {

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

func (tr *TransactionRepository) FindAllTransactions(db *sqlx.Tx, user_id int) ([]entity.Transaction, error) {
	query := `SELECT * FROM transactions WHERE user_id = $1 ORDER BY created_at DESC`

	var transactions []entity.Transaction

	err := db.Select(&transactions, query, user_id)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tr *TransactionRepository) FindTransactionById(db *sqlx.Tx, transactionId, user_id int) (*entity.Transaction, error) {

	query := `SELECT * FROM transactions WHERE id = $1 and user_id = $2`

	var transaction entity.Transaction

	err := db.Get(&transaction, query, transactionId, user_id)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (tr *TransactionRepository) UpdateTransactionById(db *sqlx.Tx, id, user_id int, reqTransaction *model.TransactionRequest) (*entity.Transaction, error) {
	query := `UPDATE transactions
SET 
	category_id = COALESCE(NULLIF($1, 0), category_id),
	amount      = COALESCE(NULLIF($2, 0), amount),
	note        = COALESCE(NULLIF($3, ''), note),
	updated_at  = now()
WHERE id = $4 AND user_id = $5
RETURNING *;
`

	var transaction entity.Transaction
	err := db.QueryRowx(query, reqTransaction.CategoryId, reqTransaction.Amount, reqTransaction.Note, id, user_id).StructScan(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil

}
