package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/starlingilcruz/golang-chat/services"
	"github.com/starlingilcruz/golang-chat/http/responses"
)

type ChatController struct {
	chatService services.Chat
}

func (c *ChatController) RegisterService(s services.Chat) {
	c.chatService = s
}

func (c *ChatController)GetMessages(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	roomId := param["id"]

	messages, _ := c.chatService.GetRoomMessages(roomId)

	d, _ := json.Marshal(responses.RoomMessagesResponse{Messages: messages})

	w.WriteHeader(http.StatusOK)
	w.Write(d)
}