package controllers

import (
	"LOCO_ASSIGNMENT/models"
	"LOCO_ASSIGNMENT/services"
	"LOCO_ASSIGNMENT/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	txnID, _ := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
    var txn models.Transaction
    if err := c.ShouldBindJSON(&txn); err != nil {
        utils.RespondError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    createdTransaction := services.CreateTransaction(txn, txnID)
    c.JSON(http.StatusCreated, createdTransaction)
}
