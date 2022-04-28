package main

import (
	"log"
	"os"

	"github.com/crafting-demo/kafka-broker-mq/pkg/websocket"
)

func main() {
	var ctx websocket.Context

	ctx.Mode = "release"
	ctx.Port = os.Getenv("KAFKA_SERVICE_PORT_API")
	if ctx.Port == "" {
		log.Fatal("KAFKA_SERVICE_PORT_API must be set")
	}

	websocket.Run(ctx)
}
