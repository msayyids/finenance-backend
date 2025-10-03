package usecase

import (
	"finenance-app/internal/entity"
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
}

func NewCategoryUseCase(db *sqlx.DB, log *zap.Logger, Validator *validator.Validate, cr repository.CategoriesRepositoryImplementation) CategoryUsecaseImplementation {
	return &CategoryUsecase{
		DB:           db,
		Log:          log,
		Validator:    Validator,
		CategoryRepo: cr,
	}
}
