package repositories

import (
	database "LOCO_ASSIGNMENT/db"
	"LOCO_ASSIGNMENT/models"
)

func CreateTransaction(txn models.Transaction, txnID int64) models.Transaction {
    return database.DB.Create(txn, txnID)
}
