package entity

import "time"

type Users struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Isverifed bool      `db:"isverifed"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
