package main

import (
	"kafka_test/config"
	"kafka_test/kafka"
	"kafka_test/modules/email_dispatcher"
	"kafka_test/modules/push_dispatcher"
	ready_process "kafka_test/modules/ready_process"
	"kafka_test/modules/recsys_queue"
)

func main() {
	cfg := config.LoadConfig()

	kafka.InitProducer(cfg.Kafka.Brokers)

	ready_process.InitReadyProcessorConsumer(cfg.Kafka.Brokers, cfg.Kafka.ConsumerGroup)
	email_dispatcher.InitEmailDispatcherConsumer(cfg.Kafka.Brokers, cfg.Kafka.ConsumerGroup)
	push_dispatcher.InitPushDispatcherConsumer(cfg.Kafka.Brokers, cfg.Kafka.ConsumerGroup)
	recsys_queue.InitRecSysQueueConsumer(cfg.Kafka.Brokers, cfg.Kafka.ConsumerGroup)

	// The application should keep running to listen for messages
	select {}
}
