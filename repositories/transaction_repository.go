package repositories

import (
	database "LOCO_ASSIGNMENT/db"
	"LOCO_ASSIGNMENT/models"
)

func CreateTransaction(txn models.Transaction, txnID int64) {
    database.DB.Create(txn, txnID)
}


func GetTransactionByID(txnID int64) (models.Transaction, error) {
    return database.DB.GetByID(txnID)
}

func GetAllTransactionsOfType(txnType string) ([]int64, error) {
    return database.DB.GetAllTransactionsOfType(txnType)
}

func GetSum(txnID int64) (float64, error) {
    return database.DB.GetSum(txnID)
}