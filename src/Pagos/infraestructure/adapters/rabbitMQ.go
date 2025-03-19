// infraestructure/adapters/rabbitMQ.go
package adapters

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"Send/src/Pagos/domain"
)


func NewRabbitMQConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://angel:1234@18.215.69.18:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to RabbitMQ")
	return conn, nil
}


type RabbitMQBroker struct {
	connection *amqp.Connection
}


func NewRabbitMQBroker(connection *amqp.Connection) domain.PagoMessageBroker {
	return &RabbitMQBroker{connection: connection}
}


func (r *RabbitMQBroker) Publish(event string, data interface{}) error {
	channel, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	// Declaramos la cola con el nombre que venga en 'event'
	_, err = channel.QueueDeclare(
		event, // Nombre de la cola
		true,  // Durable
		false, // AutoDelete
		false, // Exclusive
		false, // NoWait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Publicamos el mensaje en la cola
	err = channel.Publish(
		"",    // Exchange (por defecto)
		event, // Routing key = nombre de la cola
		false, // Mandatory
		false, // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Message sent to queue '%s': %s", event, string(body))
	return nil
}
