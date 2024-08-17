package services

import (
	"LOCO_ASSIGNMENT/models"
	"LOCO_ASSIGNMENT/repositories"
)

func CreateTransaction(txn models.Transaction, txnID int64) {
    repositories.CreateTransaction(txn, txnID)
}

func GetTransactionByID(txnID int64) (models.Transaction, error) {
    return repositories.GetTransactionByID(txnID)
}

func GetAllTransactionsOfType(txnType string) ([]int64, error) {
    return repositories.GetAllTransactionsOfType(txnType)
}

