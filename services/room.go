package services

import (
	"fmt"

	"github.com/starlingilcruz/golang-chat/internal/models"
	// "github.com/starlingilcruz/golang-chat/utils"
)


type Room struct {
	repository   models.Room
}

func (a *Room) Create(name string) (models.Room, error) {

	room := models.Room{Name: name}

	if err := room.Create(); err != nil {
		fmt.Println("Error ocurred while room creation")
	}

	return room, nil
}

func (a *Room) List() ([]models.Room, error) {
	var rooms []models.Room
	var model models.Room
	err := model.List(&rooms)
	
	if err != nil {
		fmt.Println("Error retrieving rooms")
		fmt.Println(rooms)
	}

	return rooms, nil
}

