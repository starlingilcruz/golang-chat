package rabbitmq

import (
	"os"
	"log"

	"github.com/starlingilcruz/golang-chat/services/websocket"
	"github.com/starlingilcruz/golang-chat/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type StockRequest struct {
	RoomId uint   `json:"RoomId"`
	Code   string `json:"Code"`
}

type StockResponse struct {
	RoomId  uint   `json:"RoomId"`
	Message string `json:"Message"`
}

type Broker struct {
	ConsumerQueue  amqp.Queue
	ProducerQueue  amqp.Queue
	Channel        *amqp.Channel
}

func (b *Broker) SetUp(ch *amqp.Channel) {
	consumerQueue := os.Getenv("BR_CONSUMER_QUEUE")
	producerQueue := os.Getenv("BR_PRODUCER_QUEUE")

	q1, err := ch.QueueDeclare(
		consumerQueue, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	
	q2, err := ch.QueueDeclare(
		producerQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	log.Println(err)

	b.ConsumerQueue = q1
	b.ProducerQueue = q2
	b.Channel = ch
}

// messageTransformer converts the message from rabbitmq to StockResponse type and passes it to the received messages channel
func messageTransformer(entries <-chan amqp.Delivery, receivedMessages chan StockResponse) {
	var sr StockResponse
	for d := range entries {
		err := utils.ParseByteArray(d.Body, &sr)
		if err != nil {
			log.Printf("Received bad response : %s ", string(d.Body))
			continue
		}
		log.Println("Received a response")
		receivedMessages <- sr
	}
}

// processResponse sends the stock response to the websocket's connection pool
func processResponse(s <-chan StockResponse, b *Broker, pool *websocket.Pool) {
	for r := range s {
		log.Println("processing stock response for ", r.Message)

		sr := StockResponse{
			RoomId:  r.RoomId,
			Message: r.Message,
		}

		message := websocket.Message{
			Type: 4, 
			Body: websocket.Body{
				RoomId: uint(sr.RoomId), 
				User: "stock-bot",
				Message: sr.Message,
			},
		}
		pool.Broadcast <- message
		log.Println("processed", sr.Message)
	}
}