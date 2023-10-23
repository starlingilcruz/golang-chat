package services

import (
	"log"

	"github.com/starlingilcruz/golang-chat/internal/models"
)


type Chat struct{}

type ChatRepository interface {
	GetRoomMessages(roomId string) ([]models.Chat, error)
	SaveChatMessage(msg string, roomId, userId uint) bool
}

func NewRepository() *Chat {
	return &Chat{}
}

func (c *Chat)GetRoomMessages(roomId string) ([]models.Chat, error) {
	var chats []models.Chat
	var model models.Chat
	err := model.List(roomId, &chats)
	
	if err != nil {
		log.Println("Error retrieving chats")
	}

	return chats, nil
}

func (c *Chat)SaveChatMessage(msg string, roomId, userId uint) bool {
	log.Println("--saving chat message")

	ch := models.Chat{
		Message:    msg,
		UserId:     userId,
		RoomId:     roomId,
	}

	if err := ch.Create(); err.Error != nil {
		log.Println("error: ", err.Error)
		return false
	}

	return true
}