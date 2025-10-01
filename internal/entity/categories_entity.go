package entity

import "time"

type Categories struct {
	Id         int       `db:"id"`
	Name       string    `db:"name"`
	User_Id    int       `db:"user_id"`
	Type       int       `db:"type"`
	Created_at time.Time `db:"created_at"`
}
