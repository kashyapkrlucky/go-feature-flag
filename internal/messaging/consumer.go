package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

// Consumer listens for flag change messages from RabbitMQ
type Consumer struct {
	Ch *amqp.Channel
}

// NewConsumer initializes the RabbitMQ consumer
func NewConsumer() (*Consumer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		Ch: ch,
	}, nil
}

// ListenForFlagChanges listens for flag update events from RabbitMQ
func (c *Consumer) ListenForFlagChanges() {
	msgs, err := c.Ch.Consume(
		"flag-updates", // Queue name
		"",             // Consumer name
		true,           // Auto-acknowledge
		false,          // Exclusive
		false,          // No-local
		false,          // No-wait
		nil,            // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to start consuming: %v", err)
	}

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		// Here, we can add logic to propagate changes, e.g., to other services
	}
}
