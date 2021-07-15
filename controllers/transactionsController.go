package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/AfandyW/financial-records/models"
	"github.com/bxcodec/faker/v3"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (server *Server) HomeController(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Code:    http.StatusOK,
		Data:    nil,
		Message: "Ini Halaman Home",
	}
	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (server *Server) GetAllTransactionsController(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction

	transactions, err := transaction.GetAllTransactions(server.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := Response{
		Code: http.StatusOK,
		Data: transactions,
	}

	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (server *Server) AddTransactionController(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Error read body request", http.StatusInternalServerError)
		return
	}

	timeNow := time.Now()
	transaction.Id = uuid.New().String()
	transaction.TransactionAt = timeNow.Local()
	transaction.CreateAt = transaction.TransactionAt
	transaction.UpdateAt = transaction.TransactionAt

	transactionCreated, err := transaction.CreateTransaction(server.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := Response{
		Code: http.StatusCreated,
		Data: map[string]string{
			"id":    transactionCreated.Id,
			"title": transactionCreated.Title,
		},
		Message: "New Record Transaction has been Created",
	}
	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (server *Server) GetTransactionController(w http.ResponseWriter, r *http.Request) {

	transactionId := chi.URLParam(r, "id")

	if transactionId == "" {
		http.Error(w, "Params is empty", http.StatusBadRequest)
		return
	}

	transaction := models.Transaction{}

	transactionReceive, err := transaction.GetTransaction(server.DB, transactionId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := Response{
		Code: http.StatusOK,
		Data: transactionReceive,
	}

	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (server *Server) EditTransactionController(w http.ResponseWriter, r *http.Request) {
	transactionId := chi.URLParam(r, "id")

	if transactionId == "" {
		http.Error(w, "Params is empty", http.StatusBadRequest)
		return
	}

	transaction := models.Transaction{}

	transactionReceive, err := transaction.GetTransaction(server.DB, transactionId)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	transactionUpdate := models.Transaction{}
	err = json.NewDecoder(r.Body).Decode(&transactionUpdate)

	if err != nil {
		http.Error(w, "Error read body request", http.StatusUnprocessableEntity)
		return
	}

	timeNow := time.Now()
	transactionUpdate.Id = transactionReceive.Id
	transactionUpdate.UpdateAt = timeNow.Local()

	transactionUpdated, err := transactionUpdate.UpdateTransaction(server.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := Response{
		Code: http.StatusOK,
		Data: transactionUpdated,
	}

	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (server *Server) DeleteTransactionController(w http.ResponseWriter, r *http.Request) {
	transactionId := chi.URLParam(r, "id")

	if transactionId == "" {
		http.Error(w, "Params is empty", http.StatusBadRequest)
		return
	}

	transaction := models.Transaction{}

	transactionReceive, err := transaction.GetTransaction(server.DB, transactionId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = transactionReceive.DeleteTransaction(server.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := Response{
		Code:    http.StatusOK,
		Data:    transactionId,
		Message: "Transaction has been Delete",
	}

	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (server *Server) FakeTransactionController(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}

	rand.Seed(5)

	for i := 0; i < 5; i++ {
		timeNow := time.Now()
		transaction.Id = uuid.New().String()
		transaction.Title = faker.Word()
		transaction.Description = faker.Sentence()
		transaction.TypeTransaction = faker.Word()
		transaction.Amount = rand.Intn(100)
		transaction.Currency = faker.Currency()
		transaction.Category = faker.Word()
		transaction.SubCategory = faker.Word()
		transaction.TransactionAt = timeNow.Local()
		transaction.CreateAt = transaction.TransactionAt
		transaction.UpdateAt = transaction.TransactionAt

		ttt, err := transaction.CreateTransaction(server.DB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(ttt)
	}

	res := Response{
		Code:    http.StatusCreated,
		Message: "Fake Transaction success",
	}
	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}