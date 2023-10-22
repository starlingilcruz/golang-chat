package responses

import "github.com/starlingilcruz/golang-chat/internal/models"


type RoomResponse struct {
	Room  models.Room `json:"Room"`
}

type RoomsResponse struct {
	Rooms  []models.Room `json:"Rooms"`
}