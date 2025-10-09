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
	GetAllTransactions(c fiber.Ctx) error
	GetTransactionById(c fiber.Ctx) error
	UpdateTransactionById(c fiber.Ctx) error
	DeleteTransactionById(c fiber.Ctx) error
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

func (tc *TransactionController) GetAllTransactions(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	intId, _ := strconv.Atoi(user.UserID)
	transactions, err := tc.TransactionUseCase.GetAllTransactions(intId)
	if err != nil {
		tc.Log.Error("Failed to get all transactions : %+v", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(model.WebResponse[[]entity.Transaction]{
		Code:    fiber.StatusOK,
		Message: "Transactions retrieved successfully",
		Data:    transactions,
	})
}

func (tc *TransactionController) GetTransactionById(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	intId, _ := strconv.Atoi(user.UserID)
	transactionId, _ := strconv.Atoi(c.Params("id"))

	transaction, err := tc.TransactionUseCase.GeTransactionById(transactionId, intId)
	if err != nil {
		tc.Log.Error("Failed to get transaction : %+v", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(model.WebResponse[*entity.Transaction]{
		Code:    fiber.StatusOK,
		Message: "Transaction retrieved successfully",
		Data:    transaction,
	})
}

func (tc *TransactionController) UpdateTransactionById(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	intId, _ := strconv.Atoi(user.UserID)
	transactionId, _ := strconv.Atoi(c.Params("id"))
	request := new(model.TransactionRequest)
	if err := c.Bind().Body(request); err != nil {
		tc.Log.Error("Failed to parse request body : %+v", zap.Error(err))
		return fiber.ErrBadRequest
	}

	transaction, err := tc.TransactionUseCase.UpdateTransactionById(transactionId, intId, request)
	if err != nil {
		tc.Log.Error("Failed to update transaction : %+v", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(model.WebResponse[*entity.Transaction]{
		Code:    fiber.StatusOK,
		Message: "Transaction updated successfully",
		Data:    transaction,
	})
}

func (tc *TransactionController) DeleteTransactionById(c fiber.Ctx) error {
	user := c.Locals("user").(*entity.CustomClaims)
	intId, _ := strconv.Atoi(user.UserID)
	transactionId, _ := strconv.Atoi(c.Params("id"))

	err := tc.TransactionUseCase.DeleteTransactionById(transactionId, intId)
	if err != nil {
		tc.Log.Error("Failed to delete transaction : %+v", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Transaction deleted successfully",
	})
}
