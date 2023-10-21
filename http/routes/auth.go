package routes

import (
	"github.com/gorilla/mux"

	"github.com/starlingilcruz/golang-chat/controllers"
	"github.com/starlingilcruz/golang-chat/services"
)

func RegisterAuthRoutes(router *mux.Router) {

	sr := router.PathPrefix("/v1/api/auth").Subrouter()
	// sr.Use(middlewares.HeaderMiddleware)

	var auth controllers.AuthController
	auth.RegisterService(services.Auth{})
	// auth.RegisterService(services.NewAuthService())

	sr.HandleFunc("/login", auth.Login).Methods("POST")
	sr.HandleFunc("/signup", auth.SignUp).Methods("POST")
}