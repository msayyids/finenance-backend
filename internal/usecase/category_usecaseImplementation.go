package usecase

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type CategoryUsecase struct {
	DB           *sqlx.DB
	Log          *zap.Logger
	Validator    *validator.Validate
	CategoryRepo repository.CategoriesRepositoryImplementation
}

type CategoryUsecaseImplementation interface {
	GetAllCategory(user_id int) (*[]entity.Categories, error)
	GetCategoryById(user_id, id int) (*entity.Categories, error)
	UpdateCategoryById(id, user_id int, request *model.CategoryRequest) (*entity.Categories, error)
	CreateNewCategory(user_id int, request *model.CategoryRequest) error
	DeleteCategoryById(id, user_id int) error
}

func NewCategoryUseCase(db *sqlx.DB, log *zap.Logger, Validator *validator.Validate, cr repository.CategoriesRepositoryImplementation) CategoryUsecaseImplementation {
	return &CategoryUsecase{
		DB:           db,
		Log:          log,
		Validator:    Validator,
		CategoryRepo: cr,
	}
}
