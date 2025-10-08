package config

import (
	"finenance-app/internal/delivery/http"
	"finenance-app/internal/delivery/http/middleware"
	"finenance-app/internal/delivery/http/route"
	"finenance-app/internal/repository"
	"finenance-app/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppConfig struct {
	App      *fiber.App
	DB       *sqlx.DB
	Config   *viper.Viper
	Log      *zap.Logger
	Validate *validator.Validate
	Redis    *redis.Client
}

func NewAppConfig(config *AppConfig) {
	//	setup

	categoryRepository := repository.NewCategoriesRepository()
	categoryUsecase := usecase.NewCategoryUseCase(config.DB, config.Log, config.Validate, categoryRepository)
	categoryController := http.NewCategoryController(config.Log, categoryUsecase)

	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(config.DB, config.Log, config.Validate, config.Redis, userRepository, categoryRepository)
	userController := http.NewUserController(config.Log, userUsecase)

	transactionRepository := repository.NewTransactionRepository()
	transactionUseCase := usecase.NewTransactionUseCase(config.DB, config.Log, config.Validate, transactionRepository)
	transactionController := http.NewTransactionController(config.Log, transactionUseCase)

	authMiddleware := middleware.NewAuth()

	routerConfig := route.RouteConfig{
		Logger:                config.Log,
		App:                   config.App,
		UserController:        userController,
		CategoryController:    categoryController,
		TransactionController: transactionController,
		AuthMiddleware:        authMiddleware,
	}

	routerConfig.SetupRouteConfig()
}
