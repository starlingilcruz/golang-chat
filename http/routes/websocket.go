package routes

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	// "os"

	"github.com/gorilla/mux"
	"github.com/starlingilcruz/golang-chat/services/websocket"
	"github.com/starlingilcruz/golang-chat/utils"
)

var RegisterWebSocketRoutes = func(router *mux.Router) {

	// TODO v2 - handle pool registry and to support multiple pools

	log.Println("--- configuring ws routes")
	pool := websocket.StartNewWebSocketPool()

	sb := router.PathPrefix("/v1").Subrouter()
	sb.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		token := r.URL.Query().Get("token")
		claims, _ := utils.VerifyJWT(token)

		// TODO handle unauthenticated request

		clientUser := websocket.User{
			Email:      claims["email"].(string),
			UserId:     uint(claims["userid"].(float64)),
			UserName:   claims["username"].(string),
		}

		addWsClientToPool(pool, clientUser, w, r)
	})
}

func addWsClientToPool(pool *websocket.Pool, user websocket.User, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Println(err)
	}

	client := &websocket.Client{
		Connection: conn,
		Pool:       pool,
		User:       user,
	}

	// pool.AddClient(client)
	pool.Register <- client
	requestBody := make(chan []byte)
	go client.Read(requestBody)
}