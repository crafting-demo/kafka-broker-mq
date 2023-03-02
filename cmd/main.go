package main

import (
	"os"

	"github.com/crafting-demo/kafka-broker-mq/pkg/websocket"
)

func main() {
	var ctx websocket.Context

	ctx.Mode = "release"
	ctx.Port = os.Getenv("PORT")
	if ctx.Port == "" {
		ctx.Port = "8080"
	}

	websocket.Run(ctx)
}
