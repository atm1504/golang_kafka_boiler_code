package recsys_queue

import (
	kafkaConsumer "kafka_test/kafka"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"log"
)

func InitRecSysQueueConsumer(brokers, groupID string) {
	kafkaConsumer.InitConsumer(brokers, groupID, []string{"recsys-queue"}, handleMessage)
}

func handleMessage(msg *kafka.Message) {
	log.Printf("RecSysQueue received message: %s", string(msg.Value))
	// Process message
}
