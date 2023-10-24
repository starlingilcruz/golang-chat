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
	"github.com/starlingilcruz/golang-chat/services/rabbitmq"	
)


func main() {
	fmt.Println("Starting Server...")
	
	// Load env values
	godotenv.Load()

	// Connect Rabbit MQ
	conn, ch := rabbitmq.Connect()
	defer conn.Close()
	defer ch.Close()

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
		Addr:    "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("err")
}