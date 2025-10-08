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
	CreateNewCategory(c fiber.Ctx) error
	GetAllUserCategory(c fiber.Ctx) error
	GetCategoryById(c fiber.Ctx) error
	UpdateCategoryById(c fiber.Ctx) error
	DeleteCategoryById(c fiber.Ctx) error
}

func NewCategoryController(log *zap.Logger, uc usecase.CategoryUsecaseImplementation) CategoryControllerImplementation {
	return &CategoryController{
		Log:             log,
		CategoryUsecase: uc,
	}
}

func (cc *CategoryController) CreateNewCategory(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	intId, _ := strconv.Atoi(user.UserID)

	request := new(model.CategoryRequest)
	if err := c.Bind().Body(request); err != nil {
		cc.Log.Error("Failed to parse request body : %+v", zap.Error(err))
		return fiber.ErrBadRequest
	}

	err := cc.CategoryUsecase.CreateNewCategory(intId, request)
	if err != nil {
		cc.Log.Error("Failed to create new category", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "create new category successfully",
	})
}

func (cc *CategoryController) GetAllUserCategory(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)

	intId, _ := strconv.Atoi(user.UserID)
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

func (cc *CategoryController) GetCategoryById(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	userIdint, err := strconv.Atoi(user.UserID)
	if err != nil {
		cc.Log.Error("failed to get id", zap.Error(err))
		return fiber.ErrBadRequest
	}
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		cc.Log.Error("failed to get id", zap.Error(err))
		return fiber.ErrBadRequest
	}

	category, err := cc.CategoryUsecase.GetCategoryById(userIdint, categoryId)
	if err != nil {
		cc.Log.Error("failed to get category", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(model.WebResponse[*entity.Categories]{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    category,
	})
}

func (cc *CategoryController) UpdateCategoryById(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	userIdint, err := strconv.Atoi(user.UserID)
	if err != nil {
		cc.Log.Error("failed to get id", zap.Error(err))
		return fiber.ErrBadRequest
	}
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		cc.Log.Error("failed to get id", zap.Error(err))
		return fiber.ErrBadRequest
	}

	request := new(model.CategoryRequest)
	if err := c.Bind().Body(request); err != nil {
		cc.Log.Error("Failed to parse request body : %+v", zap.Error(err))
		return fiber.ErrBadRequest
	}

	category, err := cc.CategoryUsecase.UpdateCategoryById(categoryId, userIdint, request)
	if err != nil {
		cc.Log.Error("failed to get category", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(model.WebResponse[*entity.Categories]{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    category,
	})
}
