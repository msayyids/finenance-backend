package entity

import "time"

type Categories struct {
	Id         int
	User_Id    int       `db:"user_id"`
	Name       string    `db:"name"`
	Type       string    `db:"type"`
	Created_at time.Time `db:"created_at"`
}
