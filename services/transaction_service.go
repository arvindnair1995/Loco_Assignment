package services

import (
	"LOCO_ASSIGNMENT/models"
	"LOCO_ASSIGNMENT/repositories"
)

func CreateTransaction(txn models.Transaction, txnID int64) models.Transaction {
    return repositories.CreateTransaction(txn, txnID)
}
