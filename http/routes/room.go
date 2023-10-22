package routes

import (
	"github.com/gorilla/mux"

	"github.com/starlingilcruz/golang-chat/controllers"
	"github.com/starlingilcruz/golang-chat/http/middlewares"
	"github.com/starlingilcruz/golang-chat/services"
)

func RegisterRoomRoutes(router *mux.Router) {

	sr := router.PathPrefix("/v1/api/rooms").Subrouter()
	// Add content-type json to all sub-routes
	sr.Use(middlewares.HeaderMiddleware)

	var room controllers.RoomController
	room.RegisterService(services.Room{})

	sr.HandleFunc("/", room.List).Methods("POST")
	sr.HandleFunc("/create", room.Create).Methods("POST")
}