package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "financial"
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

// var Transactions []Transaction

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func createConnection() *sql.DB {
	// Open the connection
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func main() {
	route := chi.NewRouter()

	route.Get("/", HomeController)
	route.Get("/transactions", GetAllTransactionsController)
	route.Post("/transactions", AddTransactionController)
	route.Get("/transactions/{id}", GetTransactionController)
	route.Put("/transactions/{id}", EditTransactionController)
	// route.Delete("/transactions/{id}", DeleteTransactionController)

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
	db := createConnection()
	defer db.Close()

	var transactions []Transaction

	sqlStatement := `SELECT * FROM transaction`
	rows, err := db.Query(sqlStatement)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer rows.Close()

	for rows.Next() {
		var transaction Transaction

		err = rows.Scan(&transaction.Id, &transaction.Title, &transaction.Description, &transaction.TypeTransaction, &transaction.Amount, &transaction.Currency, &transaction.Category, &transaction.SubCategory, &transaction.TransactionAt, &transaction.CreateAt, &transaction.UpdateAt)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		transactions = append(transactions, transaction)
	}

	res := Result{
		Code: http.StatusOK,
		Data: transactions,
	}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func AddTransactionController(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

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

	var id, title string

	sqlStatement := `insert into transaction (id,title,description,type_transaction,amount,currency,category,sub_category,transaction_at,create_at,update_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id, title`
	row := db.QueryRow(sqlStatement, newTransaction.Id, newTransaction.Title, newTransaction.Description, newTransaction.TypeTransaction, newTransaction.Amount, newTransaction.Currency, newTransaction.Category, newTransaction.SubCategory, newTransaction.TransactionAt, newTransaction.CreateAt, newTransaction.UpdateAt).Scan(&id, &title)

	if row != nil {
		http.Error(w, row.Error(), http.StatusInternalServerError)
	}

	// Transactions = append(Transactions, newTransaction)

	res := Result{
		Code: http.StatusCreated,
		Data: map[string]string{
			"id":    id,
			"title": title,
		},
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

func GetTransactionController(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	id := chi.URLParam(r, "id")

	sqlStatement := `SELECT * FROM transaction where id=$1`
	row, err := db.Query(sqlStatement, id)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var transaction Transaction

	err = row.Scan(&transaction.Id, &transaction.Title, &transaction.Description, &transaction.TypeTransaction, &transaction.Amount, &transaction.Currency, &transaction.Category, &transaction.SubCategory, &transaction.TransactionAt, &transaction.CreateAt, &transaction.UpdateAt)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res := Result{
		Code: http.StatusOK,
		Data: transaction,
	}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func EditTransactionController(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	id := chi.URLParam(r, "id")

	sqlStatement := `SELECT * FROM transaction where id=$1`
	_, err := db.Query(sqlStatement, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var updateTransaction Transaction

	err = json.NewDecoder(r.Body).Decode(&updateTransaction)

	if err != nil {
		http.Error(w, "Error read body request", http.StatusInternalServerError)
	}
	timeNow := time.Now()
	updateTransaction.UpdateAt = timeNow.Local()

	sqlStatement = `update transaction set title = $1,description = $2,type_transaction = $3,amount = $4,currency = $5,category = $6,sub_category = $7,update_at = $8 where id = $9`
	_, err = db.Query(sqlStatement, updateTransaction.Title, updateTransaction.Description, updateTransaction.TypeTransaction, updateTransaction.Amount, updateTransaction.Currency, updateTransaction.Category, updateTransaction.SubCategory, updateTransaction.UpdateAt, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res := Result{
		Code: http.StatusOK,
		Data: map[string]string{
			"id":    id,
			"title": updateTransaction.Title,
		},
	}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// func DeleteTransactionController(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	var index int = -1

// 	for i, x := range Transactions {
// 		if id == x.Id {
// 			index = i
// 			break
// 		}
// 	}

// 	if index == -1 {
// 		http.Error(w, "Id Not Found", http.StatusNotFound)
// 	}
// 	Transactions = append(Transactions[:index], Transactions[index+1:]...)
// 	res := Result{
// 		Code:    http.StatusOK,
// 		Message: "Record Transaction has been Delete",
// 	}
// 	result, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(result)
// }
