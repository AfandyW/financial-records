package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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
	// route.Get("/transactions/{id}", GetTransactionController)
	// route.Put("/transactions/{id}", EditTransactionController)
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

		err = rows.Scan(&transaction.Id, &transaction.Title, &transaction.Description, &transaction.TypeTransaction, &transaction.Currency, &transaction.Amount, &transaction.Category, &transaction.SubCategory, &transaction.TransactionAt)

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

	sqlStatement := `insert into transaction (id,title,description,type_transaction,amount,currency,category,sub_category,transaction_at,create_at,update_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	row, err := db.Query(sqlStatement, newTransaction.Id, newTransaction.Title, newTransaction.Description, newTransaction.TypeTransaction, newTransaction.Amount, newTransaction.Currency, newTransaction.SubCategory, newTransaction.TransactionAt, newTransaction.CreateAt, newTransaction.UpdateAt)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println(row)

	// if err != nil {
	// 	http.Error(w, "SQL Database Error", http.StatusInternalServerError)
	// }

	// Transactions = append(Transactions, newTransaction)

	// res := Result{
	// 	Code:    http.StatusCreated,
	// 	Data:    newTransaction,
	// 	Message: "New Record Transaction has been Created",
	// }
	// result, err := json.Marshal(res)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(result)
}

// func GetTransactionController(w http.ResponseWriter, r *http.Request) {
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

// 	res := Result{
// 		Code: http.StatusOK,
// 		Data: Transactions[index],
// 	}
// 	result, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(result)
// }

// func EditTransactionController(w http.ResponseWriter, r *http.Request) {
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

// 	var updateTransaction Transaction

// 	err := json.NewDecoder(r.Body).Decode(&updateTransaction)

// 	if err != nil {
// 		http.Error(w, "Error read body request", http.StatusInternalServerError)
// 	}

// 	timeNow := time.Now()

// 	Transactions[index] = Transaction{
// 		Id:              Transactions[index].Id,
// 		Title:           updateTransaction.Title,
// 		Description:     updateTransaction.Description,
// 		TypeTransaction: updateTransaction.TypeTransaction,
// 		Amount:          updateTransaction.Amount,
// 		Currency:        updateTransaction.Currency,
// 		Category:        updateTransaction.Category,
// 		SubCategory:     updateTransaction.SubCategory,
// 		UpdateAt:        timeNow.Local(),
// 		CreateAt:        Transactions[index].CreateAt,
// 		TransactionAt:   Transactions[index].TransactionAt,
// 	}

// 	res := Result{
// 		Code:    http.StatusOK,
// 		Data:    Transactions[index],
// 		Message: "New Record Transaction has been Updated",
// 	}
// 	result, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(result)
// }

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
