package kafka

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "log"
)

func ProduceMessage(topic, message string) {
    err := producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }, nil)

    if err != nil {
        log.Printf("Failed to produce message: %v", err)
    }

    // Wait for message deliveries
    go func() {
        for e := range producer.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
                } else {
                    log.Printf("Delivered message to %v\n", ev.TopicPartition)
                }
            }
        }
    }()
}
