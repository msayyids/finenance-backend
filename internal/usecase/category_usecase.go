package usecase

import (
	"errors"
	"finenance-app/internal/entity"
	"finenance-app/internal/model"

	"go.uber.org/zap"
)

func (cc *CategoryUsecase) CreateNewCategory(user_id int, request *model.CategoryRequest) error {

	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return errors.New("failed to create transaction")
	}
	defer tx.Rollback()

	if err := cc.Validator.Struct(request); err != nil {
		cc.Log.Warn("invalid user request", zap.Error(err))
		return errors.New("invalid user request")
	}

	if err := cc.CategoryRepo.AddCategory(tx, user_id, request); err != nil {
		cc.Log.Error("failed to add category", zap.Error(err))
		return errors.New("failed to add category")
	}

	err = tx.Commit()
	if err != nil {
		cc.Log.Error("failed to commit", zap.Error(err))
		return errors.New("failed to commit")
	}
	return nil
}

func (cc *CategoryUsecase) GetAllCategory(user_id int) (*[]entity.Categories, error) {

	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return &[]entity.Categories{}, errors.New("failed to create transaction")
	}
	defer tx.Rollback()

	category, err := cc.CategoryRepo.GetAllCategories(tx, user_id)
	if err != nil {
		cc.Log.Error("failed to get category", zap.Error(err))
		return &[]entity.Categories{}, errors.New("failed to get category")
	}
	err = tx.Commit()
	if err != nil {
		cc.Log.Error("failed to commit", zap.Error(err))
		return &[]entity.Categories{}, errors.New("failed to commit")
	}

	// final response

	return category, nil
}

func (cc *CategoryUsecase) GetCategoryById(user_id, id int) (*entity.Categories, error) {

	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to create transaction")
	}

	defer tx.Rollback()

	category, err := cc.CategoryRepo.GetCategoriesById(tx, user_id, id)
	if err != nil {
		cc.Log.Error("failed to get category", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to get category")
	}
	err = tx.Commit()
	if err != nil {
		cc.Log.Error("failed to commit", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to commit")
	}

	// final response

	return category, nil
}
func (cc *CategoryUsecase) UpdateCategoryById(id, user_id int, request *model.CategoryRequest) (*entity.Categories, error) {
	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to create transaction")
	}
	defer tx.Rollback()

	if err := cc.Validator.Struct(request); err != nil {
		cc.Log.Warn("invalid user request", zap.Error(err))
		return &entity.Categories{}, errors.New("invalid user request")
	}

	updatedCategory, err := cc.CategoryRepo.UpdateCategoryById(tx, id, user_id, request)
	if err != nil {
		cc.Log.Error("failed to update category", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to update category")
	}

	if err := tx.Commit(); err != nil {
		cc.Log.Error("failed to commit", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to commit")
	}

	return updatedCategory, nil
}
