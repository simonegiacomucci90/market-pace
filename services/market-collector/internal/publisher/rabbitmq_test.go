package publisher

import (
	"os"
	"testing"
)

func TestRabbitMQPublish(t *testing.T) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		t.Skip("RABBITMQ_URL not set, skipping test")
	}

	publisher, err := NewRabbitMQPublisher(url, "test_queue")
	if err != nil {
		t.Fatalf("NewRabbitMQPublisher failed: %v", err)
	}

	data := []byte("test data")
	err = publisher.RabbitMQPublish(data)
	if err != nil {
		t.Fatalf("RabbitMQPublish failed: %v", err)
	}
}
