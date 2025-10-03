package http

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type CategoryController struct {
	Log             *zap.Logger
	CategoryUsecase usecase.CategoryUsecaseImplementation
}

type CategoryControllerImplementation interface {
	GetAllUserCategory(c fiber.Ctx) error
}

func NewCategoryController(log *zap.Logger, uc usecase.CategoryUsecaseImplementation) CategoryControllerImplementation {
	return &CategoryController{
		Log:             log,
		CategoryUsecase: uc,
	}
}

func (cc *CategoryController) GetAllUserCategory(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)

	intId, err := strconv.Atoi(user.UserID)
	categoryUser, err := cc.CategoryUsecase.GetAllCategory(intId)
	if err != nil {
		cc.Log.Error("failed to get all category", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(model.WebResponse[*[]entity.Categories]{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    categoryUser,
	})
}
