package usecase

import (
	"finenance-app/internal/model"
	"finenance-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserUsecaseImplementation interface {
	Create(request *model.UsersRegisterRequest) (model.UsersResponseRegister, error)
	Login(request *model.UserLoginRequest) (model.UserLoginResponse, error)
	GetProfile(userId string) (model.UserGetProfileResponse, error)
	Logout(refreshToken string) error
}
type UserUsecase struct {
	DB           *sqlx.DB
	Log          *zap.Logger
	Validator    *validator.Validate
	ReddisClient *redis.Client
	UserRepo     repository.UserRepositoryImplementation
}

func NewUserUsecase(db *sqlx.DB, log *zap.Logger, v *validator.Validate, r *redis.Client, ur repository.UserRepositoryImplementation) UserUsecaseImplementation {
	return &UserUsecase{
		DB:           db,
		Log:          log,
		Validator:    v,
		ReddisClient: r,
		UserRepo:     ur,
	}
}
