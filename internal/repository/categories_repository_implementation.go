package repository

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"

	"github.com/jmoiron/sqlx"
)

type CategoriesRepository struct{}

type CategoriesRepositoryImplementation interface {
	AddCategory(db *sqlx.Tx, user_id int, request *model.CategoryRequest) error
	GetAllCategories(db *sqlx.Tx, user_id int) (*[]entity.Categories, error)
	AddDefaultCategory(db *sqlx.Tx, categories []entity.Categories) error
	GetCategoriesById(db *sqlx.Tx, userId int, id int) (*entity.Categories, error)
	UpdateCategoryById(db *sqlx.Tx, id, user_id int, request *model.CategoryRequest) (*entity.Categories, error)
	DeleteCategoryById(db *sqlx.Tx, id, user_id int) error
}

func NewCategoriesRepository() CategoriesRepositoryImplementation {
	return &CategoriesRepository{}
}
