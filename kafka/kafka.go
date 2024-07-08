package kafka

import (
	"kafka_test/config"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	producer *kafka.Producer
	consumer *kafka.Consumer
)

func InitKafka(cfg config.Config) {
	var err error

	// Initialize producer with idempotence enabled
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.Kafka.Brokers,
		"enable.idempotence": true,
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	// Initialize consumer
	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Brokers,
		"group.id":          cfg.Kafka.ConsumerGroup,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
}
