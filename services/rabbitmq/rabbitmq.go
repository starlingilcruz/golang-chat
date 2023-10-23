package rabbitmq

import (
	"log"
	"fmt"
	"os"
	amqp "github.com/rabbitmq/amqp091-go"

)

var br Broker


func Connect() (*amqp.Connection, *amqp.Channel) {
	amqHost := os.Getenv("RMQ_HOST")
	amqUser := os.Getenv("RMQ_USERNAME")
	amqPass := os.Getenv("RMQ_PASSWORD")
	amqPort := os.Getenv("RMQ_PORT")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", amqUser, amqPass, amqHost, amqPort))

	if err != nil {
		log.Println("Error: could not initialize RabbitMQ")
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Println("Error: could not create channel on RabbitMQ")
	}

	br.SetUp(ch)

	return conn, ch
}

func GetRabbitMQBroker() *Broker {
	return &br
}