package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

func (cr *CategoriesRepository) AddCategory(db *sqlx.Tx, category *entity.Categories) error {
	query := `INSERT INTO categories (user_id, name,type) VALUES ($1, $2, $3) RETURNING id, created_at`

	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowx(category).StructScan(category)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoriesRepository) GetAllCategories(db *sqlx.Tx, userId int) (*[]entity.Categories, error) {
	var categories []entity.Categories

	query := `SELECT * FROM categories WHERE user_id = $1`

	err := db.QueryRowx(query, userId).StructScan(&categories)
	if err != nil {
		return nil, err
	}

	return &categories, nil

}

func (cr *CategoriesRepository) GetCategoriesById(db *sqlx.Tx, userId int, id int) (*entity.Categories, error) {
	var categories entity.Categories

	query := `SELECT * FROM categories WHERE user_id = :$1 and id = :$2`

	err := db.QueryRowx(query, userId, id).StructScan(&categories)
	if err != nil {
		return nil, err
	}

	return &categories, nil

}
