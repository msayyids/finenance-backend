package http

import (
	"finenance-app/internal/entity"
	"finenance-app/internal/model"
	"finenance-app/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type TransactionController struct {
	Log                *zap.Logger
	TransactionUseCase usecase.TransactionUseCaseImpl
}

type TransactionControllerImpl interface {
	CreateNewTransaction(c fiber.Ctx) error
}

func NewTransactionController(log *zap.Logger, transactionUseCase usecase.TransactionUseCaseImpl) TransactionControllerImpl {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
		Log:                log,
	}
}

func (tc *TransactionController) CreateNewTransaction(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	intId, _ := strconv.Atoi(user.UserID)

	request := new(model.TransactionRequest)
	if err := c.Bind().Body(request); err != nil {
		tc.Log.Error("Failed to parse request body : %+v", zap.Error(err))
		return fiber.ErrBadRequest
	}

	err := tc.TransactionUseCase.CreateTransaction(intId, request)
	if err != nil {
		tc.Log.Error("Failed to create transaction : %+v", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "Transaction created successfully",
	})

}
