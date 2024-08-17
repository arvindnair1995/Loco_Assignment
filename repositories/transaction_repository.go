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