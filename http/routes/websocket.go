package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	gws "github.com/gorilla/websocket"

	"github.com/starlingilcruz/golang-chat/services/websocket"
	"github.com/starlingilcruz/golang-chat/services/rabbitmq"
	"github.com/starlingilcruz/golang-chat/utils"
)

type Tenant struct {
	Pool   *websocket.Pool
	Broker *rabbitmq.Broker
}

type TenantRegistry struct {
	Room   map[string]*Tenant
}

var RegisterWebSocketRoutes = func(router *mux.Router) {

	tenantRegistry := TenantRegistry{
		Room:   make(map[string]*Tenant),
	}

	log.Println("--- configuring ws routes")

	sb := router.PathPrefix("/v1").Subrouter()
	sb.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		roomId    := r.URL.Query().Get("roomId")
		token     := r.URL.Query().Get("token")
		claims, _ := utils.VerifyJWT(token)
		
		// TODO handle unauthenticated request
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

		addWsClientToPool(conn, &tenantRegistry, roomId, clientUser)
	})
}

func addWsClientToPool(conn *gws.Conn, registry *TenantRegistry, roomId string, user websocket.User) {

	tenant := registry.Room[roomId]

	if tenant == nil {
		pool := websocket.StartNewWebSocketPool()
		br   := rabbitmq.GetRabbitMQBroker()

		tenant = &Tenant{
			Pool:    pool,
			Broker:  br,
		}
		registry.Room[roomId] = tenant
	}

	rabbitmq.GetRabbitMQBroker()

	pool := tenant.Pool
	br   := tenant.Broker

	client := &websocket.Client{
		Connection: conn,
		Pool:       pool,
		User:       user,
	}

	pool.AddClient(client)
	bodyChannel := make(chan []byte)
	go client.Read(bodyChannel)
	go br.ReadMessages(pool)
	go br.PublishMessage(bodyChannel)
}