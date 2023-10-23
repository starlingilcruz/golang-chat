package rabbitmq

import (
	"log"

	"github.com/starlingilcruz/golang-chat/services/websocket"
)

func (b *Broker) ReadMessages(pool *websocket.Pool) {
	msgs, err := b.Channel.Consume(
		b.ConsumerQueue.Name, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	
	if err != nil {
		log.Printf("ReadMessages error occured %s\n", err)
		return
	}

	rsvdMsgs := make(chan StockResponse)
	go messageTransformer(msgs, rsvdMsgs)
	go processResponse(rsvdMsgs, b, pool)
	select {}
}