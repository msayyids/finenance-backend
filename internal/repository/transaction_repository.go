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

func (tr *TransactiomRepository) FindTransactionById(db *sqlx.DB, transactionId int) (*entity.Transaction, error) {

	query := `SELECT * FROM transactions WHERE id = $1`

	var transaction entity.Transaction

	err := db.Get(&transaction, query, transactionId)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (tr *TransactiomRepository) UpdateTransaction(db *sqlx.Tx, transaction *entity.Transaction) (*entity.Transaction, error) {
	query := `UPDATE transactions 
	SET category_id  = COALESCE(NULLIF($1, ''), category_id), amount = COALESCE(NULLIF($2, ''), amount),
	note = COALESCE(NULLIF($3, ''), note),updated_at   = now()
    WHERE id = $4 and user_id = $5
	RETURNING id, user_id, category_id, amount, note, created_at,updated_at;`

	err := db.QueryRowx(
		query,
		transaction.CategoryId,
		transaction.Amount,
		transaction.Note,
		transaction.UpdatedAt,
		transaction.Id,
		transaction.UserId,
	).StructScan(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil

}

func (tr *TransactiomRepository) DeleteTransactionById(db *sqlx.DB, transactionId, user_id int) error {
	query := `DELETE FROM transactions WHERE id = $1 and user_id = $2;`
	_, err := db.Exec(query, transactionId, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TransactiomRepository) FindTransactionsByCategory(db *sqlx.DB, categoryId, user_id int) (*[]entity.Transaction, error) {
	query := `SELECT category_id,amount,note FROM transactions WHERE category_id = $1 and user_id = $2;`

	var transactions []entity.Transaction
	err := db.Select(&transactions, query, categoryId, user_id)
	if err != nil {
		return nil, err
	}
	return &transactions, nil

}

func (tr *TransactiomRepository) FindTransactionsByType(db *sqlx.DB, categoryId, user_id int, types string) (*[]entity.Transaction, error) {
	query := `select c.name ,c.type,amount,note from transactions
	inner join public.categories c on transactions.category_id = c.id
	where c.type =$1 and c.user_id = $2`

	var transactions []entity.Transaction
	err := db.Select(&transactions, query, categoryId, user_id)
	if err != nil {
		return nil, err
	}
	return &transactions, nil
}
