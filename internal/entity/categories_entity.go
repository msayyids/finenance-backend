package entity

import "time"

type Categories struct {
	Id         int       `db:"id"`
	User_Id    int       `db:"user_id"`
	Name       string    `db:"name"`
	Type       string    `db:"type"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}
