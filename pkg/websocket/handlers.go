package websocket

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/crafting-demo/kafka-broker-mq/pkg/kafka"
	"github.com/crafting-demo/kafka-broker-mq/pkg/logger"
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
		logger.Write("ProducerHandler", "failed to upgrade request", err)
		return
	}
	defer ws.Close()

	var msg Message
	if err := ws.ReadJSON(&msg); err != nil {
		logger.Write("ProducerHandler", "failed to read json message", err)
		return
	}

	m, err := json.Marshal(msg)
	if err != nil {
		logger.Write("ProducerHandler", "failed to marshal json message", err)
		return
	}

	var producer kafka.Producer
	if err := producer.Enqueue(c.Param("topic"), m); err != nil {
		logger.Write("ProducerHandler", "failed to enqueue message", err)
	}
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
		logger.Write("ConsumerHandler", "failed to upgrade request", err)
		return
	}
	defer ws.Close()

	var consumer kafka.Consumer

	conn, err := consumer.New()
	if err != nil {
		logger.Write("ConsumerHandler", "failed to create new consumer", err)
		return
	}
	defer conn.Close()

	var saramaOffset int64
	offset := c.Param("offset")
	if offset == "latest" {
		saramaOffset = sarama.OffsetNewest
	} else {
		saramaOffset = sarama.OffsetOldest
	}

	partitionConsumer, err := conn.ConsumePartition(c.Param("topic"), 0, saramaOffset)
	if err != nil {
		logger.Write("ConsumerHandler", "failed to create partition consumer", err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var message Message
			if err := json.Unmarshal(msg.Value, &message); err != nil {
				logger.Write("ConsumerHandler", "failed to parse json encoded message", err)
				continue
			}
			if err := ws.WriteJSON(message); err != nil {
				logger.Write("ConsumerHandler", "failed to write json", err)
			}
		case <-signals:
			return
		}
	}
}
