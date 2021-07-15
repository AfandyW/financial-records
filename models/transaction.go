package models

import (
	"database/sql"
	"fmt"
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

func (t *Transaction) CreateTransaction(db *sql.DB) (*Transaction, error) {

	sqlStatement := `INSERT INTO transaction (
		id,title,description,type_transaction,
		amount,currency,category,sub_category,
		transaction_at,create_at,update_at) 
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	_, err := db.Query(
		sqlStatement,
		t.Id,
		t.Title,
		t.Description,
		t.TypeTransaction,
		t.Amount,
		t.Currency,
		t.Category,
		t.SubCategory,
		t.TransactionAt,
		t.CreateAt,
		t.UpdateAt)

	if err != nil {
		return &Transaction{}, err
	}

	return t, nil
}

func (t *Transaction) GetAllTransactions(db *sql.DB) (*[]Transaction, error) {

	transactions := []Transaction{}

	sqlStatement := `SELECT * FROM transaction`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		return &[]Transaction{}, err
	}

	for rows.Next() {
		var transaction Transaction
		err = rows.Scan(
			&transaction.Id,
			&transaction.Title,
			&transaction.Description,
			&transaction.TypeTransaction,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Category,
			&transaction.SubCategory,
			&transaction.TransactionAt,
			&transaction.CreateAt,
			&transaction.UpdateAt)

		if err != nil {
			return &[]Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}

	return &transactions, nil
}

func (t *Transaction) GetTransaction(db *sql.DB, transactionId string) (*Transaction, error) {

	transaction := Transaction{}
	sqlStatement := `SELECT * FROM transaction where id=$1`
	row := db.QueryRow(sqlStatement, transactionId)

	err := row.Scan(
		&transaction.Id,
		&transaction.Title,
		&transaction.Description,
		&transaction.TypeTransaction,
		&transaction.Amount,
		&transaction.Currency,
		&transaction.Category,
		&transaction.SubCategory,
		&transaction.TransactionAt,
		&transaction.CreateAt,
		&transaction.UpdateAt)

	if err != nil {
		return &Transaction{}, err
	}

	return &transaction, nil
}

func (t *Transaction) UpdateTransaction(db *sql.DB) (*Transaction, error) {

	sqlStatement := `UPDATE transaction SET 
	title = $1,
	description = $2,
	type_transaction = $3,
	amount = $4,
	currency = $5,
	category = $6,
	sub_category = $7,
	update_at = $8 WHERE id = $9`

	_, err := db.Query(
		sqlStatement,
		t.Title,
		t.Description,
		t.TypeTransaction,
		t.Amount,
		t.Currency,
		t.Category,
		t.SubCategory,
		t.UpdateAt,
		t.Id)

	if err != nil {
		fmt.Println(err)
		return &Transaction{}, err
	}

	return t, nil
}

func (t *Transaction) DeleteTransaction(db *sql.DB) (*Transaction, error) {

	sqlStatement := `delete from transaction where id = $1`
	_, err := db.Query(sqlStatement, t.Id)

	if err != nil {
		return &Transaction{}, err
	}

	return t, nil
}
