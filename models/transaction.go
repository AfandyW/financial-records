package models

import (
	"time"
)

type Transaction struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	TypeTransaction string    `json:"type_transaction"`
	Amount          int       `json:"amount"`
	Currency        string    `json:"currency"`
	Category        string    `json:"category"`
	SubCategory     string    `json:"sub_category"`
	TransactionAt   time.Time `json:"transaction_at"`
	CreateAt        time.Time `json:"create_at"`
	UpdateAt        time.Time `json:"update_at"`
}
