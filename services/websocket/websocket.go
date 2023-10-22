package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	
	return conn, err
}



