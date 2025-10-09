package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct{}

type UserRepositoryImplementation interface {
	AddUser(db *sqlx.Tx, user *entity.Users) error
	FindUserById(db *sqlx.Tx, userid int) (*entity.Users, error)
	FindUserByEmail(db *sqlx.Tx, email string) (*entity.Users, error)
}

func NewUserRepository() UserRepositoryImplementation {
	return &UserRepository{}
}
