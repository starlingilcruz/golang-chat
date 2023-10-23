package websocket


import (
	"fmt"
	"encoding/json"
	"strings"

	"github.com/gorilla/websocket"
	// "github.com/starlingilcruz/golang-chat/services"
)

type Client struct {
	ID         string
	Connection *websocket.Conn
	Pool       *Pool
	User       User
}

type User struct {
	Email    string  `json:"email,omitempty"`
	UserId   uint    `json:"userId,omitempty"`
	UserName string  `json:"userName,omitempty"`
}

type Message struct {
	Type int `json:"Type,omitempty"`
	Body Body
}

type Body struct {
	RoomName string `json:"roomName,omitempty"`
	RoomId   uint  `json:"roomId,omitempty"`
	Message  string `json:"message,omitempty"`
	User     string `json:"user,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (c *Client) Read(channel chan []byte) {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()
	defer c.Pool.ReviveWebsocket()

	// var repo services.ChatRepository = services.NewRepository()

	for {
		mType, payload, err := c.Connection.ReadMessage()
		
		var body Body
		err = json.Unmarshal(payload, &body)

		if err != nil {
			fmt.Println(err)
		}

		body.User = c.User.UserName
		body.Email = c.User.Email
		message := Message{Type: mType, Body: body}
		c.Pool.Broadcast <- message
		
		if strings.Index(body.Message, "/stock=") == 0 {
			channel <- payload
		} else {
			// repo.SaveChatMessage(body.Message, uint(body.RoomId), c.User.UserId)
		}
	}
}
