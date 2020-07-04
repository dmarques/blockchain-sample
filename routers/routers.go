package routers

import (
	"net/http"

	"github.com/blockchain-sample/controllers"
	"github.com/gin-gonic/gin"
)

func SetAllRouters(router *gin.Engine) {

	SetPingRouter(router)
	SetNewBlock(router)
	SetGetBlockchain(router)
}

func SetPingRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func SetNewBlock(router *gin.Engine) {
	router.POST("/block", controllers.CreateNewBlock)
}

func SetGetBlockchain(router *gin.Engine) {
	router.GET("/chain", controllers.GetBlockchain)
}
