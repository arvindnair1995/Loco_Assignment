package database

import (
	"LOCO_ASSIGNMENT/models"
	"LOCO_ASSIGNMENT/utils"
	"container/list"
	"errors"
)

type InMemoryDB struct {
    store map[int64]models.Transaction
    typeMap map[string][]int64
    graph map[int64][]int64
}

var DB = &InMemoryDB{
    store: make(map[int64]models.Transaction),
    typeMap: make(map[string][]int64),
    graph: make(map[int64][]int64),
}

func (db *InMemoryDB) Create(txn models.Transaction, txnID int64) {
    db.store[txnID] = txn
    db.typeMap[utils.TransformString(txn.Type)] = append(db.typeMap[txn.Type], txnID)
    if txn.ParentID != nil {
        db.graph[*txn.ParentID] = append(db.graph[*txn.ParentID], txnID)
    }
}

func (db *InMemoryDB) GetByID(txnID int64) (models.Transaction, error) {
    txn, exists := db.store[txnID]
    if !exists {
        return models.Transaction{}, errors.New("transaction not found")
    }
    return txn, nil
}

func (db *InMemoryDB) GetAllTransactionsOfType(txnType string) ([]int64, error) {
    txnIDs, exists := db.typeMap[txnType]
    if !exists {
        return nil, errors.New("transaction type not found")
    }
    return txnIDs, nil
}

func (db *InMemoryDB) GetSum(txnID int64) (float64, error) {
    queue := list.New()
    queue.PushBack(txnID)

    sum := 0.0

    for queue.Len() > 0 {
        element := queue.Front()
        current := element.Value.(int64)
        queue.Remove(element)

        sum += db.store[current].Amount

        neighbors, exists := db.graph[current]
        if exists {
            for _, neighbor := range neighbors {
                queue.PushBack(neighbor)
            }
        }
    }
    return sum, nil
}
