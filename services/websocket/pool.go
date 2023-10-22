package websocket

import (
	"log"
	"os"
	"runtime/debug"

)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	defer p.ReviveWebsocket()

	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			log.Println("info:", "New client. Size of connection pool:", len(p.Clients))
			for c := range p.Clients {
				err := c.Connection.WriteJSON(Message{Type: 2, Body: Body{Message: "new user joined..."}})
				log.Println(err)
			}

		case client := <-p.Unregister:
			// remove client from pool
			delete(p.Clients, client)
			log.Println("info:", "disconnected a client. size of connection pool:", len(p.Clients))
			for c := range p.Clients {
				err := c.Connection.WriteJSON(Message{Type: 3, Body: Body{Message: "user disconnected..."}})
				log.Println(err)
			}

		case msg := <-p.Broadcast:
			log.Println("info", "broadcast message to clients in pool")
			for c := range p.Clients {
				err := c.Connection.WriteJSON(msg)
				log.Println(err)
			}
		}
	}
}

func StartNewWebSocketPool() *Pool {
	log.Println("--- starting ws pool")

	pool := NewPool()
	go pool.Start()

	return pool
}

func (p *Pool) AddClient(client *Client) {
	p.Register <- client
}

func (p *Pool) ReviveWebsocket() {
	if err := recover(); err != nil {
		if os.Getenv("LOG_PANIC_TRACE") == "true" {
			log.Println(
				"level:", "error",
				"err: ", err,
				"trace", string(debug.Stack()),
			)
		} else {
			log.Println(
				"level", "error",
				"err", err,
			)
		}
		go p.Start()
	}
}