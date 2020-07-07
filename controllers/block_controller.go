package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blockchain-sample/utils"
)

func GetBlockchain(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ReadChainFileAndParse())
}
