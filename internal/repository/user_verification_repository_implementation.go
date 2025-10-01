package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserVerificationRepository struct {
}

type UserVerificationRepositoryImplementation interface {
	CreateVerificationKey(db *sqlx.Tx, userVerification *entity.UserVerification) (*entity.UserVerification, error)
	FindUserVerifByKey(db *sqlx.Tx, userVerification *entity.UserVerification) (*entity.UserVerification, error)
}

func NewUserVerification() UserVerificationRepositoryImplementation {
	return &UserVerificationRepository{}
}
