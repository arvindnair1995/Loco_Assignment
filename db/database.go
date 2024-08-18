package database

import (
	"LOCO_ASSIGNMENT/models"
	"LOCO_ASSIGNMENT/utils"
	"container/list"
	"errors"
	"sync"
)

type InMemoryDB struct {
    store map[int64]models.Transaction
    typeMap map[string][]int64
    graph map[int64][]int64
    mu sync.Mutex
}

var DB = &InMemoryDB{
    store: make(map[int64]models.Transaction),
    typeMap: make(map[string][]int64),
    graph: make(map[int64][]int64),
}

// TC: O(1)
func (db *InMemoryDB) Create(txn models.Transaction, txnID int64) {
    db.mu.Lock()
    defer db.mu.Unlock()

    db.store[txnID] = txn
    txn.Type = utils.TransformString(txn.Type)
    db.typeMap[txn.Type] = append(db.typeMap[txn.Type], txnID)
    if txn.ParentID != nil {
        db.graph[*txn.ParentID] = append(db.graph[*txn.ParentID], txnID)
    }
}

// TC: O(1)
func (db *InMemoryDB) GetByID(txnID int64) (models.Transaction, error) {
    db.mu.Lock()
    defer db.mu.Unlock()

    txn, exists := db.store[txnID]
    if !exists {
        return models.Transaction{}, errors.New("transaction not found")
    }
    return txn, nil
}

// TC: O(1)
func (db *InMemoryDB) GetAllTransactionsOfType(txnType string) ([]int64, error) {
    db.mu.Lock()
    defer db.mu.Unlock()

    txnIDs, exists := db.typeMap[txnType]
    if !exists {
        return nil, errors.New("transaction type not found")
    }
    return txnIDs, nil
}

// TC: O(n)
// BFS to compute sum
func (db *InMemoryDB) GetSum(txnID int64) (float64, error) {
    db.mu.Lock()
    defer db.mu.Unlock()
    
    _, exists := db.store[txnID]
    if !exists {
        return 0.0, errors.New("transaction not found")
    }
    
    visited := make(map[int64]bool)
    queue := list.New()

    queue.PushBack(txnID)
    visited[txnID] = true

    sum := 0.0

    for queue.Len() > 0 {
        ele := queue.Front()
        curr := ele.Value.(int64)
        queue.Remove(ele)

        sum += db.store[curr].Amount

        children, exists := db.graph[curr]
        if exists {
            for _, child := range children {
                if !visited[child] {
                    queue.PushBack(child)
                    visited[child] = true
                }
            }
        }
    }
    return sum, nil
}
