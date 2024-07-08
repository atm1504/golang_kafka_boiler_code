package ready_processor

import (
	kafkaConsumer "kafka_test/kafka"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func InitReadyProcessorConsumer(brokers, groupID string) {
	kafkaConsumer.InitConsumer(brokers, groupID, []string{"testing-ready-queue"}, handleMessage)
}

func handleMessage(msg *kafka.Message) {
	log.Printf("ReadyProcessor received message: %s", string(msg.Value))
	// Process message
}
