package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

func (vr *UserVerificationRepository) CreateVerificationKey(db *sqlx.Tx, userVerification *entity.UserVerification) (*entity.UserVerification, error) {
	query := `INSERT INTO verification_keys (user_id, key, expired_at) 
              VALUES (:user_id, :key, :expired_at) 
              RETURNING id`

	err := db.QueryRowx(
		query,
		userVerification.UserId,
		userVerification.Key,
		userVerification.ExpiredAt,
	).Scan(&userVerification.Id)

	if err != nil {
		return nil, err
	}

	return userVerification, nil
}

func (vr *UserVerificationRepository) FindUserVerifByKey(db *sqlx.Tx, userVerification *entity.UserVerification) (*entity.UserVerification, error) {
	query := `
        SELECT id, user_id, key, expired_at
        FROM verification_keys
        WHERE user_id = $1 AND key = $2 AND expired_at > NOW()
        LIMIT 1
    `

	err := db.Get(userVerification, query, userVerification.UserId, userVerification.Key)
	if err != nil {
		return nil, err
	}

	return userVerification, nil
}

func (vr *UserVerificationRepository) UpdateVerificationByUserId(db *sqlx.Tx) (*entity.UserVerification, error) {

	return &entity.UserVerification{}, nil
}
