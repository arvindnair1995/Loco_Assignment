package database

import (
	"LOCO_ASSIGNMENT/models"
	"errors"
)

type InMemoryDB struct {
    store map[int64]models.Transaction
    typeMap map[string][]models.Transaction
}

var DB = &InMemoryDB{
    store: make(map[int64]models.Transaction),
    typeMap: make(map[string][]models.Transaction),
}

func (db *InMemoryDB) Create(txn models.Transaction, txnID int64) models.Transaction {
    db.store[txnID] = txn
    db.typeMap[txn.Type] = append(db.typeMap[txn.Type], txn)
    return txn
}

func (db *InMemoryDB) GetByID(txnID int64) (models.Transaction, error) {
    txn, exists := db.store[txnID]
    if !exists {
        return models.Transaction{}, errors.New("transaction not found")
    }
    return txn, nil
}
