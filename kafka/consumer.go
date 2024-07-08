package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func InitConsumer(brokers, groupID string, topics []string, handler func(*kafka.Message)) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"group.id":           groupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %s", err)
	}

	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				handler(msg)
			} else {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()
}
