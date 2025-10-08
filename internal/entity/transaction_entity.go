package entity

import "time"

type Transaction struct {
	id         int       `db:"id"`
	UserId     int       `db:"user_id"`
	CategoryId int       `db:"category_id"`
	Amount     int       `db:"amount"`
	Note       string    `db:"note"`
	CreateAt   time.Time `db:"created_at"`
	UpdateAt   time.Time `db:"updated_at"`
}
