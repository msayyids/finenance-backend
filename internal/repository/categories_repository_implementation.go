package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

type CategoriesRepository struct{}

type CategoriesRepositoryImplementation interface {
	AddCategory(db *sqlx.Tx, category *entity.Categories) error
	GetAllCategories(db *sqlx.Tx, user_id int) (*[]entity.Categories, error)
	AddDefaultCategory(db *sqlx.Tx, categories []entity.Categories) error
}

func NewCategoriesRepository() CategoriesRepositoryImplementation {
	return &CategoriesRepository{}
}
