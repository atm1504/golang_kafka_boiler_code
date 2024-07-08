package email_dispatcher

import (
	kafkaConsumer "kafka_test/kafka"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func InitEmailDispatcherConsumer(brokers, groupID string) {
	kafkaConsumer.InitConsumer(brokers, groupID, []string{"email-dispatcher"}, handleMessage)
}

func handleMessage(msg *kafka.Message) {
	log.Printf("EmailDispatcher received message: %s", string(msg.Value))
	// Process message
}
