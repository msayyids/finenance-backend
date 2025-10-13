package usecase

import (
	"errors"
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/utils"

	"go.uber.org/zap"
)

func (cc *CategoryUsecase) CreateNewCategory(user_id int, request *model.CategoryRequest) error {

	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return errors.New("failed to create transaction")
	}
	defer utils.CommitOrRollback(tx, &err)

	if err := cc.Validator.Struct(request); err != nil {
		cc.Log.Warn("invalid user request", zap.Error(err))
		return errors.New("invalid user request")
	}

	if err := cc.CategoryRepo.AddCategory(tx, user_id, request); err != nil {
		cc.Log.Error("failed to add category", zap.Error(err))
		return errors.New("failed to add category")
	}

	return nil
}

func (cc *CategoryUsecase) GetAllCategory(user_id int) (*[]entity.Categories, error) {

	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return &[]entity.Categories{}, errors.New("failed to create transaction")
	}
	defer utils.CommitOrRollback(tx, &err)

	category, err := cc.CategoryRepo.GetAllCategories(tx, user_id)
	if err != nil {
		cc.Log.Error("failed to get category", zap.Error(err))
		return &[]entity.Categories{}, errors.New("failed to get category")
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

	defer utils.CommitOrRollback(tx, &err)

	category, err := cc.CategoryRepo.GetCategoriesById(tx, user_id, id)
	if err != nil {
		cc.Log.Error("failed to get category", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to get category")
	}

	// final response

	return category, nil
}
func (cc *CategoryUsecase) UpdateCategoryById(id, user_id int, request *model.CategoryRequest) (
	*entity.Categories, error,
) {
	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to create transaction")
	}
	defer utils.CommitOrRollback(tx, &err)

	if err := cc.Validator.Struct(request); err != nil {
		cc.Log.Warn("invalid user request", zap.Error(err))
		return &entity.Categories{}, errors.New("invalid user request")
	}

	updatedCategory, err := cc.CategoryRepo.UpdateCategoryById(tx, id, user_id, request)
	if err != nil {
		cc.Log.Error("failed to update category", zap.Error(err))
		return &entity.Categories{}, errors.New("failed to update category")
	}

	return updatedCategory, nil
}

func (cc *CategoryUsecase) DeleteCategoryById(user_id, id int) error {
	tx, err := cc.DB.Beginx()
	if err != nil {
		cc.Log.Error("failed to create transaction", zap.Error(err))
		return errors.New("failed to create transaction")
	}

	defer utils.CommitOrRollback(tx, &err)
	if err := cc.CategoryRepo.DeleteCategoryById(tx, user_id, id); err != nil {
		cc.Log.Error("failed to delete category", zap.Error(err))
		return errors.New("failed to delete category")
	}
	
	return nil
}
