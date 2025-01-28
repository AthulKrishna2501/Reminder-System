package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var Channel *amqp091.Channel
var Connection *amqp091.Connection

func Init() error {
	var err error
	Connection, err = amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}

	Channel, err = Connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}

	_, err = Channel.QueueDeclare(
		"reminder_queue", // Queue name
		true,             // Durable (survives server restarts)
		false,            // Auto-delete
		false,            // Exclusive
		false,            // No-wait
		nil,              // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return err
	}

	return nil
}

func Close() {
	if Channel != nil {
		Channel.Close()
	}
	if Connection != nil {
		Connection.Close()
	}
}
