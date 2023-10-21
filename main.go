package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/starlingilcruz/golang-chat/internal/db"
	"github.com/starlingilcruz/golang-chat/internal/models"
)


func main() {
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
}