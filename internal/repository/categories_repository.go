package repository

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (cr *CategoriesRepository) AddCategory(db *sqlx.Tx, user_id int, request *model.CategoryRequest) error {
	query := `INSERT INTO categories (user_id, name,type) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, user_id, request.Name, request.Type)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoriesRepository) AddDefaultCategory(db *sqlx.Tx, categories []entity.Categories) error {
	query := `
INSERT INTO categories (user_id, name, type)
VALUES (:user_id, :name, :type)
`

	_, err := db.NamedExec(query, categories)
	if err != nil {
		return fmt.Errorf("failed to insert categories: %w", err)
	}

	return nil
}

func (cr *CategoriesRepository) GetAllCategories(db *sqlx.Tx, userId int) (*[]entity.Categories, error) {
	var categories []entity.Categories

	query := `SELECT * FROM categories WHERE user_id = $1`

	err := db.Select(&categories, query, userId)
	if err != nil {
		return nil, err
	}

	return &categories, nil

}

func (cr *CategoriesRepository) GetCategoriesById(db *sqlx.Tx, userId int, id int) (*entity.Categories, error) {
	var categories entity.Categories

	query := `SELECT * FROM categories WHERE user_id = $1 and id = $2`

	err := db.QueryRowx(query, userId, id).StructScan(&categories)
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (cr *CategoriesRepository) UpdateCategoryById(db *sqlx.Tx, id, user_id int, request *model.CategoryRequest) (*entity.Categories, error) {

	query := `UPDATE categories 
	SET name  = COALESCE(NULLIF($1, ''), name), type = COALESCE(NULLIF($2, ''), type), updated_at   = now()
    WHERE id = $3 and user_id = $4
	RETURNING id, user_id, name, type, created_at, updated_at;
`
	var updatedCategory entity.Categories
	err := db.QueryRowx(query, request.Name, request.Type, id, user_id).StructScan(&updatedCategory)
	if err != nil {
		return nil, err
	}

	return &updatedCategory, nil
}

func (cr *CategoriesRepository) DeleteCategoryById(db *sqlx.Tx, id, user_id int) error {
	query := `DELETE FROM categories WHERE id = $1 and user_id = $2`
	_, err := db.Exec(query, id, user_id)
	if err != nil {
		return err
	}
	return nil
}
