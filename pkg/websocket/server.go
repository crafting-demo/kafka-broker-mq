package websocket

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	Mode string
	Port string
}

func Run(ctx Context) {
	gin.SetMode(ctx.Mode)

	router := gin.Default()

	router.GET("/producer/:topic", ProducerHandler)
	router.GET("/consumer/:topic/:offset", ConsumerHandler)

	router.Run(":" + ctx.Port)
}
