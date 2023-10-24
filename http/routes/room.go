package routes

import (
	"log"
	"github.com/gorilla/mux"

	"github.com/starlingilcruz/golang-chat/controllers"
	"github.com/starlingilcruz/golang-chat/http/middlewares"
	"github.com/starlingilcruz/golang-chat/services"
)

func RegisterRoomRoutes(router *mux.Router) {

	log.Println("--- configuring room http routes")

	sr := router.PathPrefix("/v1/api/rooms").Subrouter()
	// Add content-type json to all sub-routes
	sr.Use(middlewares.HeaderMiddleware)
	sr.Use(middlewares.AuthMiddleware)

	var room controllers.RoomController
	room.RegisterService(services.Room{})

	var chat controllers.ChatController
	chat.RegisterService(services.Chat{})

	sr.HandleFunc("/", room.List).Methods("POST")
	sr.HandleFunc("/create", room.Create).Methods("POST")
	sr.HandleFunc("/{id}/messages", chat.GetMessages).Methods("GET")

}