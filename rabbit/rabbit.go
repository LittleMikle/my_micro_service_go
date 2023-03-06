package rabbit

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbit() (err error) {
	var conn *amqp091.Connection
	conn, err = amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("failed with connection to Rabbit: %w", err)
	}

	var ch *amqp091.Channel
	ch, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("failed with opening a channel: %w", err)
	}
	defer ch.Close()
	return nil
}
