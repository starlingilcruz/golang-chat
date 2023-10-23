package rabbitmq

import (
	"log"
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (b *Broker) PublishMessage(requestBody chan []byte) {
	for body := range requestBody {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		err := b.Channel.PublishWithContext(ctx,
			"",                    // exchange
			b.ProducerQueue.Name, // routing key
			false,                 // mandatory
			false,                 // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		cancel()
		if err != nil {
			log.Printf("PublishMessage Error occured %s\n", err)
			continue
		}
		log.Printf(" [x] Sent %s\n", body)
	}
}