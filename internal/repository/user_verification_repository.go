package repository

import (
	"finenance-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

func (vr *UserVerificationRepository) CreateVerificationKey(db *sqlx.Tx) (*entity.UserVerification, error) {
	query := `INSERT INTO verification_keys (user_id,key,expired_at,created_at)  VALUES (:user_id,:key,:expired_at,)`

	if err:=db.


	return nil
}
