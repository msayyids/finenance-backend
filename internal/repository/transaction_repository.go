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

func (tr *TransactionRepository) FindTransactionById(db *sqlx.Tx, transactionId, user_id int) (
	*entity.Transaction, error,
) {

	query := `SELECT * FROM transactions WHERE id = $1 and user_id = $2`

	var transaction entity.Transaction

	err := db.Get(&transaction, query, transactionId, user_id)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (tr *TransactionRepository) UpdateTransactionById(
	db *sqlx.Tx, id, user_id int, reqTransaction *model.TransactionRequest,
) (*entity.Transaction, error) {
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
	err := db.QueryRowx(
		query, reqTransaction.CategoryId, reqTransaction.Amount, reqTransaction.Note, id, user_id,
	).StructScan(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil

}

func (tr *TransactionRepository) DeleteTransactionById(db *sqlx.Tx, id, user_id int) error {
	query := `DELETE FROM transactions WHERE id = $1 AND user_id = $2`
	_, err := db.Exec(query, id, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TransactionRepository) FindTransactionByIncome(db *sqlx.Tx, userId int) (
	*[]model.TransactionTypeResponse, error,
) {
	query := `
SELECT 
    t.id,
    t.amount,
    t.note,
    c.name AS category_name,
    c.type,
    t.created_at
FROM transactions t
JOIN categories c ON c.id = t.category_id
WHERE t.user_id = $1 AND LOWER(c.type) = LOWER('income');
	`

	var transactions []model.TransactionTypeResponse
	err := db.Select(&transactions, query, userId)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (tr *TransactionRepository) FindTransactionByExpense(db *sqlx.Tx, userId int) (
	*[]model.TransactionTypeResponse, error,
) {
	query := `
SELECT 
    t.id,
    t.amount,
    t.note,
    c.name AS category_name,
    c.type,
    t.created_at
FROM transactions t
JOIN categories c ON c.id = t.category_id
WHERE t.user_id = $1 AND LOWER(c.type) = LOWER('expense');
	`

	var transactions []model.TransactionTypeResponse
	err := db.Select(&transactions, query, userId)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}
