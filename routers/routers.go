package routers

import (
	"net/http"

	"github.com/blockchain-sample/controllers"
	"github.com/gin-gonic/gin"
)

func SetAllRouters(router *gin.Engine) {

	SetPingRouter(router)
	SetTransactionRouter(router)
	SetBlockchainRouter(router)
}

func SetPingRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func SetTransactionRouter(router *gin.Engine) {
	router.GET("/pending_transactions", controllers.GetPendingTransactions)

	router.POST("/transaction", controllers.CreateTransaction)
	router.POST("/mine", controllers.ProcessPendingTransactions)
}

func SetBlockchainRouter(router *gin.Engine) {
	router.GET("/chain", controllers.GetBlockchain)

}
