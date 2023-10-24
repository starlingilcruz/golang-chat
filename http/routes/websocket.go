package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/starlingilcruz/golang-chat/services/websocket"
	"github.com/starlingilcruz/golang-chat/services/rabbitmq"
	"github.com/starlingilcruz/golang-chat/utils"
)

type Channel struct {
	Pool   *websocket.Pool
	Broker *rabbitmq.Broker
}

type ChannelRegistry struct {
	Room   map[string]*Channel
}

var RegisterWebSocketRoutes = func(router *mux.Router) {

	log.Println("--- configuring ws routes")

	channelRegistry := ChannelRegistry{
		Room:   make(map[string]*Channel),
	}

	sb := router.PathPrefix("/v1").Subrouter()
	sb.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		roomId    := r.URL.Query().Get("roomId")
		token     := r.URL.Query().Get("token")
		claims, _ := utils.VerifyJWT(token)

		if utils.IsValidToken(token) == false {
			return
		}
		
		clientUser := websocket.User{
			Email:      claims["email"].(string),
			UserId:     uint(claims["userid"].(float64)),
			UserName:   claims["username"].(string),
		}

		conn, err := websocket.Upgrade(w, r)

		if err != nil {
			log.Println(err)
			return
		}

		broadcast := make(chan []byte)
		pool, _   := registerChannelDeps(&channelRegistry, roomId, broadcast)

		client := &websocket.Client{
			Connection: conn,
			Pool:       pool,
			User:       clientUser,
		}
	
		pool.AddClient(client)
		go client.Read(broadcast)
	})
}

func registerChannelDeps(registry *ChannelRegistry, roomId string, broadcast chan []byte) (*websocket.Pool, *rabbitmq.Broker) {
	channel := registry.Room[roomId]

	if channel == nil {
		channel = &Channel{
			Pool:    websocket.StartNewWebSocketPool(),
			Broker:  rabbitmq.GetRabbitMQBroker(),
		}
		registry.Room[roomId] = channel
	}

	pool := channel.Pool
	br   := channel.Broker

	go br.ReadMessages(pool)
	go br.PublishMessage(broadcast)

	return pool, br
}	