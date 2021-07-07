package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

var Transactions []Transaction

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	route := chi.NewRouter()

	route.Get("/", HomeController)
	route.Get("/transactions", GetAllTransactionsController)
	route.Post("/transactions", AddTransactionsController)
	// route.Post("/transactions/{id}", EditTransactionsController)

	http.ListenAndServe(":3000", route)

}

func HomeController(w http.ResponseWriter, r *http.Request) {
	res := Result{
		Code:    http.StatusOK,
		Data:    nil,
		Message: "Ini Halaman Home",
	}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetAllTransactionsController(w http.ResponseWriter, r *http.Request) {
	res := Result{
		Code:    http.StatusOK,
		Data:    Transactions,
		Message: nil,
	}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func AddTransactionsController(w http.ResponseWriter, r *http.Request) {
	var newTransaction Transaction

	err := json.NewDecoder(r.Body).Decode(&newTransaction)

	if err != nil {
		http.Error(w, "Error read body request", http.StatusInternalServerError)
	}
	timeNow := time.Now()
	newTransaction.Id = uuid.New().String()
	newTransaction.TransactionAt = timeNow.Local()
	newTransaction.CreateAt = newTransaction.TransactionAt
	newTransaction.UpdateAt = newTransaction.TransactionAt

	Transactions = append(Transactions, newTransaction)

	res := Result{
		Code:    http.StatusCreated,
		Data:    Transactions,
		Message: "New Record Transaction has been Created",
	}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}
