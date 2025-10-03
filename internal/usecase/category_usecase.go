package usecase

import (
	"errors"
	"finenance-app/internal/entity"

	"go.uber.org/zap"
)

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
