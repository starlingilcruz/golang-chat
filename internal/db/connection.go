package connection

import (
	"os"
	"fmt"

	"gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSL_MODE")
	timeZone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", 
		host, dbUser, dbPwd, dbName, port, sslMode, timeZone
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}