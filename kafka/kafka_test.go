package kafka

import (
	"kafka_test/config"
	"testing"
)

func TestKafka(t *testing.T) {
	cfg := config.LoadConfig()
	InitKafka(cfg)

	ProduceMessage("test_topic", "test message")

	// More tests for consumer can be added
}
