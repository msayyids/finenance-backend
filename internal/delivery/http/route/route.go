package route

import (
	"finenance-app/internal/delivery/http"
	"finenance-app/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type RouteConfig struct {
	Logger                *zap.Logger
	App                   *fiber.App
	UserController        http.UserControllerImplementation
	CategoryController    http.CategoryControllerImplementation
	TransactionController http.TransactionControllerImpl
	AuthMiddleware        fiber.Handler
	LoggerMiddleware      fiber.Handler
}

func (c *RouteConfig) SetupRouteConfig() {
	c.App.Use(middleware.LogMiddleware(c.Logger))
	c.SetupGlobalRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGlobalRoute() {
	c.App.Post("/finenance/signup", c.UserController.Register)
	c.App.Post("/finenance/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(middleware.NewAuth())
	c.App.Get("finenance/user", c.UserController.GetProfile)
	c.App.Delete("/finenance/logout", c.UserController.Logout)

	c.App.Post("finenance/category", c.CategoryController.CreateNewCategory)
	c.App.Get("finenance/category", c.CategoryController.GetAllUserCategory)
	c.App.Get("finenance/category/:id", c.CategoryController.GetCategoryById)
	c.App.Patch("finenance/category/:id", c.CategoryController.UpdateCategoryById)

	c.App.Post("finenance/transaction", c.TransactionController.CreateNewTransaction)
	c.App.Get("finenance/transaction/", c.TransactionController.GetAllTransactions)
	c.App.Get("finenance/transaction/:id", c.TransactionController.GetTransactionById)
	c.App.Patch("finenance/transaction/:id", c.TransactionController.UpdateTransactionById)
	c.App.Delete("finenance/transaction/:id", c.TransactionController.DeleteTransactionById)
	c.App.Get("finenance/transaction/type/income", c.TransactionController.GetTransactionsByIncome)
	c.App.Get("finenance/transaction/type/expense", c.TransactionController.GetTransactionsByExpense)
}
