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
	userVerifRepo := repository.NewUserVerificationRepository()
	userVerifUsecase := usecase.NewUserVerifUsecase(config.DB, config.Log, config.Validate, config.Redis, userVerifRepo)

	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(config.DB, config.Log, config.Validate, config.Redis, userRepository)
	userController := http.NewUserController(config.Log, userUsecase, userVerifUsecase)

	authMiddleware := middleware.NewAuth()

	routerconfig := route.RouteConfig{
		Logger:         config.Log,
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}

	routerconfig.SetupRouteConfig()
}
