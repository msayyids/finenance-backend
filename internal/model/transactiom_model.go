package model

type TransactionRequest struct {
	CategoryId int    `json:"category_id"`
	Amount     int    `json:"amount"`
	Note       string `json:"note"`
}
