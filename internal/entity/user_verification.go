package entity

import "time"

type UserVerification struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	Key       string    `db:"key"`
	ExpiredAt time.Time `db:"expired_at"`
	IsUsed    bool      `db:"is_used"`
	createdAt time.Time `db:"created_at"`
}
