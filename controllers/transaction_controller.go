package controllers

import (
	"net/http"

	"github.com/blockchain-sample/domain"
	"github.com/blockchain-sample/services"
	"github.com/blockchain-sample/utils"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {

	var dataTransaction domain.Transaction

	//Receive JSON file
	if errBind := c.BindJSON(&dataTransaction); errBind != nil {
		c.JSON(http.StatusBadRequest, "Invalid Json format.")
		return
	}

	//Add new Transaction
	services.AddNewTransaction(dataTransaction)

	c.JSON(http.StatusOK, "Transaction was created.")
}

func GetPendingTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ReadPendingTransactionFile())
}

func ProcessPendingTransactions(c *gin.Context) {
	services.ProcessPendingTransactions()
	c.JSON(http.StatusOK, "Pending transactions processed.")
}
