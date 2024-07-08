package push_dispatcher

import (
	kafkaConsumer "kafka_test/kafka"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func InitPushDispatcherConsumer(brokers, groupID string) {
	kafkaConsumer.InitConsumer(brokers, groupID, []string{"push-dispatcher"}, handleMessage)
}

func handleMessage(msg *kafka.Message) {
	log.Printf("PushDispatcher received message: %s", string(msg.Value))
	// Process message
}
