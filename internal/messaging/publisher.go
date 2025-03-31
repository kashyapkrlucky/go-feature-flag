package messaging

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Publisher struct holds the connection details
type Publisher struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

// NewPublisher initializes the RabbitMQ connection and channel
func NewPublisher() (*Publisher, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &Publisher{
		Conn: conn,
		Ch:   ch,
	}, nil
}

// PublishFlagChange sends a message to the "flag-updates" queue
func (p *Publisher) PublishFlagChange(flagID int64, action string) error {
	message := fmt.Sprintf("Feature flag %d %s", flagID, action)
	err := p.Ch.Publish(
		"",             // Default exchange
		"flag-updates", // Queue name
		false,          // Mandatory
		false,          // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}
	log.Printf("Message published: %s", message)
	return nil
}
