package websocket

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/crafting-demo/kafka-broker-mq/pkg/kafka"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ProducerHandler(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Upgrade request to websocket protocol.
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	var msg kafka.Message
	if err := ws.ReadJSON(&msg); err != nil {
		return
	}

	producer := kafka.Producer{Topic: c.Param("topic")}
	producer.Enqueue(msg)
}

func ConsumerHandler(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Upgrade request to websocket protocol.
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	msgCh := make(chan []byte)
	doneCh := make(chan struct{}, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	consumer := kafka.Consumer{Topic: c.Param("topic")}
	go consumer.Run(msgCh, doneCh)

	for {
		select {
		case <-ctx.Done():
			return
		case <-doneCh:
			return
		case m := <-msgCh:
			var msg kafka.Message
			if err := json.Unmarshal(m, &msg); err == nil {
				ws.WriteJSON(msg)
			}
		}
	}
}
