package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

func (r *UserRepository) AddUser(db *sqlx.Tx, user *entity.Users) error {

	query := `
INSERT INTO users (name, email, password, created_at, updated_at)
VALUES (:name, :email, :password, NOW(), NOW())
RETURNING id, created_at, updated_at
`

	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowx(user).StructScan(user)
	if err != nil {
		return err
	}

	return nil

}

func (r *UserRepository) FindUserById(db *sqlx.Tx, userId int) (*entity.Users, error) {
	query := `SELECT * FROM users WHERE id=$1`

	var user entity.Users
	err := db.QueryRowx(query, userId).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindUserByEmail(db *sqlx.Tx, email string) (*entity.Users, error) {
	query := `SELECT * FROM users WHERE email=$1`

	var user entity.Users
	err := db.QueryRowx(query, email).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser() error { return nil }

func (r *UserRepository) DeleteUser() error { return nil }

func (r *UserRepository) FindRole() error { return nil }
