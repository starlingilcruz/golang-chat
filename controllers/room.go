package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"


	"github.com/starlingilcruz/golang-chat/services"
	"github.com/starlingilcruz/golang-chat/utils"
	"github.com/starlingilcruz/golang-chat/http/requests"
	"github.com/starlingilcruz/golang-chat/http/responses"
)

type RoomController struct {
	// TODO dependency injection
	roomService services.Room
}

func (rc *RoomController) RegisterService(s services.Room) {
	rc.roomService = s
}

func (rc *RoomController)List(w http.ResponseWriter, r *http.Request) {
		rooms, _ := rc.roomService.List()

		d, _ := json.Marshal(responses.RoomsResponse{Rooms: rooms})

		w.WriteHeader(http.StatusOK)
		w.Write(d)
}

func (rc *RoomController)Create(w http.ResponseWriter, r *http.Request) {
	var params requests.RoomCreateParams

	if err := utils.ParseBody(r, &params); err != nil {
		fmt.Println("Error parsing request payload")
		return
	}

	room, _ := rc.roomService.Create(params.Name)

	d, _ := json.Marshal(responses.RoomResponse{Room: room})

	w.WriteHeader(http.StatusOK)
	w.Write(d)
}
