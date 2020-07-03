package main

import (
	"github.com/blockchain-sample/routers"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func init() {
	routers.SetAllRouters(Router)
}

func main() {

	Router.Run(":8080")

}
