package websocket

import "log"

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
			log.Println("WS: new client. Pool size:", len(p.Clients))
			for c := range p.Clients {
				err := c.Connection.WriteJSON(
					Message{
						Type: 2, 
						Body: Body{Message: "new user joined..."},
					},
				)
				log.Println(err)
			}

		case client := <-p.Unregister:
			delete(p.Clients, client)
			log.Println("WS: client disconnected. Pool size: ", len(p.Clients))
		
			for c := range p.Clients {
				err := c.Connection.WriteJSON(
					Message{
						Type: 3, 
						Body: 
						Body{Message: "user disconnected..."},
					},
				)
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
		log.Println("Revive websocket: ", err)
		go p.Start()
	}
}