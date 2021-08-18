package repository

import (
	"database/sql"
	"fmt"

	"github.com/AfandyW/financial-records/models"
)

var orderTransaction string = `id,title,description,type_transaction,amount,currency,category,sub_category,transaction_at,create_at,update_at`

type transactionRepositoryImpl struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepositoryImpl{DB: db}
}

func (repository *transactionRepositoryImpl) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	db := repository.DB

	sqlStatement := `INSERT INTO transactions (
		` + orderTransaction + `) 
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	row, err := db.Query(
		sqlStatement,
		transaction.Id,
		transaction.Title,
		transaction.Description,
		transaction.TypeTransaction,
		transaction.Amount,
		transaction.Currency,
		transaction.Category,
		transaction.SubCategory,
		transaction.TransactionAt,
		transaction.CreateAt,
		transaction.UpdateAt)

	defer row.Close()

	if err != nil {
		return &models.Transaction{}, err
	}

	return transaction, nil
}

func (repository *transactionRepositoryImpl) GetAllTransactions(limit int, page int) ([]models.Transaction, error) {
	db := repository.DB
	transactions := []models.Transaction{}

	var err error
	var rows *sql.Rows

	if limit != 0 {
		if page != 0 {
			offset := (page - 1) * limit
			sqlStatement := `SELECT ` + orderTransaction + ` FROM transactions LIMIT $1 OFFSET $2`
			rows, err = db.Query(sqlStatement, limit, offset)
		} else {
			sqlStatement := `SELECT ` + orderTransaction + ` FROM transactions LIMIT $1`
			rows, err = db.Query(sqlStatement, limit)
		}
	} else {
		sqlStatement := `SELECT ` + orderTransaction + ` FROM transactions`
		rows, err = db.Query(sqlStatement)
	}

	if err != nil {
		return []models.Transaction{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
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
			return []models.Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (repository *transactionRepositoryImpl) GetTransaction(transactionId string) (*models.Transaction, error) {
	db := repository.DB
	transaction := models.Transaction{}
	sqlStatement := `SELECT ` + orderTransaction + `
	FROM transactions where id=$1`
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
		return &models.Transaction{}, err
	}

	return &transaction, nil
}

func (repository *transactionRepositoryImpl) UpdateTransaction(transactionId string, transaction *models.Transaction) (*models.Transaction, error) {
	db := repository.DB
	sqlStatement := `UPDATE transactions SET 
	title = $1,
	description = $2,
	type_transaction = $3,
	amount = $4,
	currency = $5,
	category = $6,
	sub_category = $7,
	update_at = $8 WHERE id = $9`

	row, err := db.Query(
		sqlStatement,
		transaction.Title,
		transaction.Description,
		transaction.TypeTransaction,
		transaction.Amount,
		transaction.Currency,
		transaction.Category,
		transaction.SubCategory,
		transaction.UpdateAt,
		transaction.Id)

	defer row.Close()

	if err != nil {
		fmt.Println(err)
		return &models.Transaction{}, err
	}

	return transaction, nil
}

func (repository *transactionRepositoryImpl) DeleteTransaction(transactionId string) (bool, error) {
	db := repository.DB
	sqlStatement := `delete from transactions where id = $1`
	row, err := db.Query(sqlStatement, transactionId)

	row.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}
