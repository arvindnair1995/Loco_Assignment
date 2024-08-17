package controllers

import (
	"LOCO_ASSIGNMENT/models"
	"LOCO_ASSIGNMENT/services"
	"LOCO_ASSIGNMENT/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTransaction godoc
// @Summary Create a transaction
// @Description Create a transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction_id path int64 true "Transaction ID"
// @Param txn body models.Transaction true "Transaction Data"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transaction/{transaction_id} [put]
func CreateTransaction(c *gin.Context) {
	txnID, _ := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	var txn models.Transaction
	if err := c.ShouldBindJSON(&txn); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	services.CreateTransaction(txn, txnID)
	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

// GetTransaction godoc
// @Summary Get a transaction
// @Description Get a transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction_id path int64 true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 500 {object} map[string]string
// @Router /transaction/{transaction_id} [get]
func GetTransactionByID(c *gin.Context) {
	txnID, _ := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	txn, err := services.GetTransactionByID(txnID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Transaction not found")
		return
	}
	c.JSON(http.StatusOK, txn)
}

// GetTransactionsOfType godoc
// @Summary Get all transactions of a given type
// @Description Get all transactions of a given type
// @Tags transactions
// @Accept json
// @Produce json
// @Param type path string true "Transaction type"
// @Success 200 {object} []int64
// @Failure 500 {object} map[string]string
// @Router /types/{type} [get]
func GetAllTransactionsOfType(c *gin.Context) {
	txnType := c.Param("type")
	txnIDs, err := services.GetAllTransactionsOfType(txnType)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Transaction type not found")
		return
	}
	c.JSON(http.StatusOK, txnIDs)
}

// GetSumOfTransaction godoc
// @Summary Get sum of transaction
// @Description Get sum of transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction_id path int64 true "Transaction ID"
// @Success 200 {object} map[string]float64 
// @Failure 500 {object} map[string]string
// @Router /sum/{transaction_id} [get]
func GetSum(c *gin.Context) {
	txnID, _ := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	sum, err := services.GetSum(txnID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error encountered")
		return
	}
	c.JSON(http.StatusOK, gin.H{"sum": sum})
}