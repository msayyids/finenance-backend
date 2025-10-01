package usecase

import (
	"finenance-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserVerificationUsecase struct {
	DB            *sqlx.DB
	Log           *zap.Logger
	Validator     *validator.Validate
	ReddisClient  *redis.Client
	UserVerifRepo repository.UserVerificationRepositoryImplementation
}

type UserVerifUsecaseImplementation interface {
}

func NewUserVerifUsecase(db *sqlx.DB, log *zap.Logger, v *validator.Validate, r *redis.Client, uv repository.UserVerificationRepositoryImplementation) UserVerifUsecaseImplementation {
	return &UserVerificationUsecase{
		DB:            db,
		Log:           log,
		Validator:     v,
		ReddisClient:  r,
		UserVerifRepo: uv,
	}
}
