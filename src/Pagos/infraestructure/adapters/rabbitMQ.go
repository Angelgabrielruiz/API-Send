package adapters

import (
	"github.com/streadway/amqp"
	"log"
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
