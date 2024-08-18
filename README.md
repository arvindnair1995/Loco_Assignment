# Loco Backend Assignment

Go + Gin + Swagger project involving CRUD APIs for a transaction service. All the transactions are stored in a custom in-memory database. This DB has 3 data structures - 
1. `Store `- For storing txnID to txn mapping
2. `TypeMap` - For storing txnType to txn mapping
3. `Graph` - For storing dependencies between various transactions

## Setup
1. Git clone the project
2. Do `go mod download` to install all the dependencies
3. Open [Swagger UI](http://localhost:8080/swagger/index.html#/) on your browser
4. Voila! You are already to use the transaction service

## APIs
1. `PUT /transactionservice/transaction/{transaction_id}` - Inserts a transaction into the DB with TC: O(1)
2. `GET /transactionservice/transaction/{transaction_id}` - Returns the details of the particular txn. TC: O(1) as it's just returning the value from a hashmap
3. `GET /transactionservice/types/{type}` - Returns a list of txnIDs who have the same type. TC: O(1) due to hashmap retrieval
4. `GET /transactionservice/sum/{transaction_id}` - Returns the sum of all trnsactions linked to the transaction_id. TC: O(n) where n is the number of transactions. BFS has been used to compute the sum.