package config

import (
	"os"
)

type Config struct {
	Kafka struct {
		Brokers       string
		ConsumerGroup string
	}
}

func LoadConfig() Config {
	var cfg Config
	cfg.Kafka.Brokers = os.Getenv("KAFKA_BROKERS")
	cfg.Kafka.ConsumerGroup = os.Getenv("KAFKA_CONSUMER_GROUP")
	return cfg
}
