package application

import (
	"log"
	"github.com/streadway/amqp"
)

type RabbitRepository interface {
	SendMessage(queueName string, message string) error
}

type rabbitRepository struct {
	connection *amqp.Connection
}

func NewRabbitRepository(connection *amqp.Connection) RabbitRepository {
	return &rabbitRepository{connection: connection}
}

func (r *rabbitRepository) SendMessage(queueName string, message string) error {
	channel, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	// Declare a queue
	_, err = channel.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // AutoDelete
		false,     // Exclusive
		false,     // NoWait
		nil,       // Arguments
	)
	if err != nil {
		return err
	}

	// Publish a message to the queue
	err = channel.Publish(
		"",        // Exchange
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Message sent to queue: %s", message)
	return nil
}
