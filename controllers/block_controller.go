package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blockchain-sample/config"
	"github.com/blockchain-sample/domain"
	"github.com/blockchain-sample/services"
)

func CreateNewBlock(c *gin.Context) {

	var dataTransaction domain.DataTransaction

	//Receive JSON file
	if errBind := c.BindJSON(&dataTransaction); errBind != nil {
		c.JSON(http.StatusBadRequest, "Invalid Json format.")
		return
	}

	//Create new Block
	block, err := services.CreateBlockServices(dataTransaction.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
	}

	//Add new Block
	// services.AddNewBlock()

	c.JSON(http.StatusOK, block.Hash)
}

func GetBlockchain(c *gin.Context) {

	c.JSON(http.StatusOK, config.ReadFileAndParse())
}