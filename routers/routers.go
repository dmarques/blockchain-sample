package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAllRouters(router *gin.Engine) {

	SetPingRouter(router)
}

func SetPingRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
