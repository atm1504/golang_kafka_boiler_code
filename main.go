package main

import (
	"kafka_test/config"
	"kafka_test/kafka"
)

func main() {
	// Initialize config
	cfg := config.LoadConfig()

	// Initialize Kafka
	kafka.InitKafka(cfg)

	// Example usage
	kafka.ProduceMessage("my_topic", "Hello Kafka")

	// Start consuming messages
	go kafka.StartConsumer("my_topic", "my_group")

	// Keep the application running
	select {}
}
