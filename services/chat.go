package services

import (
	"log"

	"github.com/starlingilcruz/golang-chat/internal/models"
)


type chat struct{}

type ChatRepository interface {
	SaveChatMessage(msg string, roomId, userId uint) bool
}

func NewRepository() *chat {
	return &chat{}
}

func (c *chat) SaveChatMessage(msg string, roomId, userId uint) bool {
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