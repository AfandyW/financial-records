package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Transaction struct {
	Id              uint   `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	TypeTransaction string `json:"type_transaction"`
	Amount          int    `json:"amount"`
	Currency        string `json:"currency"`
	Category        string `json:"category"`
	SubCategory     string `json:"sub_category"`
	TransactionAt   string `json:"transaction_at"`
	CreateAt        string `json:"create_at"`
	UpdateAt        string `json:"update_at"`
}

var Transactions []Transaction

func main() {
	route := chi.NewRouter()

	route.Get("/", HomeController)
	route.Get("/transactions", GetAllTransactionsController)

	http.ListenAndServe(":3000", route)

}

func HomeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hay, Welcome"))
}

func GetAllTransactionsController(w http.ResponseWriter, r *http.Request) {

}
