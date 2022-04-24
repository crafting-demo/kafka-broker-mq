package websocket

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"

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

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	consumer := kafka.Consumer{Topic: c.Param("topic")}

	messages, err := consumer.Messages()
	if err != nil {
		return
	}

	for {
		select {
		case <-signalCh:
			return
		case msg := <-messages:
			var message kafka.Message
			if err := json.Unmarshal(msg.Value, &message); err == nil {
				ws.WriteJSON(msg)
			}
		}
	}
}
