package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"

	"github.com/starlingilcruz/golang-chat/internal/db"
	"github.com/starlingilcruz/golang-chat/internal/models"
	"github.com/starlingilcruz/golang-chat/http/routes"
	"github.com/starlingilcruz/golang-chat/http/middlewares"

	
)


func main() {
	fmt.Println("Starting Server...")
	
	// Load env values
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// instantiate mux routes
	db.Connect()
	dbi := db.GetInstance()
	dbi.AutoMigrate(models.Tables...)
	fmt.Println("Database tables migrated")

	StartHttpServer()
}

func StartHttpServer() {
	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)
	routes.RegisterRoomRoutes(r)
	routes.RegisterWebSocketRoutes(r)

	handler := middlewares.CORS(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("err")
}