package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
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
		log.Println("Failed to upgrade request", err)
		return
	}
	defer ws.Close()

	var msg kafka.Message
	if err := ws.ReadJSON(&msg); err != nil {
		log.Println("Failed to read json", err)
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
		log.Println("Failed to upgrade request", err)
		return
	}
	defer ws.Close()

	consumer := kafka.Consumer{Topic: c.Param("topic")}

	conn, err := consumer.New()
	if err != nil {
		log.Println("Failed to create new consumer", err)
		return
	}
	defer conn.Close()

	partitionConsumer, err := conn.ConsumePartition(consumer.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Println("Failed to create partition consumer", err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var message kafka.Message
			if err := json.Unmarshal(msg.Value, &message); err != nil {
				log.Println("Failed to parse json encoded message", err)
				ws.WriteJSON(msg.Value)
				continue
			}
			ws.WriteJSON(message)
		case <-signals:
			return
		}
	}
}
