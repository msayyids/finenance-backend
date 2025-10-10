package model

import "time"

type TransactionRequest struct {
	CategoryId *int    `json:"category_id"`
	Amount     *int    `json:"amount"`
	Note       *string `json:"note"`
}

type TransactionTypeResponse struct {
	ID           int       `db:"id" json:"id"`
	Amount       float64   `db:"amount" json:"amount"`
	Note         string    `db:"note" json:"note"`
	CategoryName string    `db:"category_name" json:"category_name"`
	Type         string    `db:"type" json:"type"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
