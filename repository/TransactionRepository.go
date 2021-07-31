package repository

import (
	"github.com/AfandyW/financial-records/models"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	GetAllTransactions(limit int, page int) (*[]models.Transaction, error)
	GetTransaction(transactionId string) (*models.Transaction, error)
	UpdateTransaction(transactionId string, transaction *models.Transaction) (*models.Transaction, error)
	DeleteTransaction(transactionId string) (bool, error)
}
