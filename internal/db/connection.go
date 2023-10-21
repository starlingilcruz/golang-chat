package db

import (
	"os"
	"fmt"

	"gorm.io/driver/postgres"
  "gorm.io/gorm"
)

// type Database struct {
// 	Instance *gorm.DB
// }

var (
	db *gorm.DB
)

func Connect() error {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSL_MODE")
	timeZone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, dbUser, dbPwd, dbName, port, sslMode, timeZone)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Has ocurred an error during database connection")
	}

	db = d

	fmt.Println("Database connected!")

	return err
}

func GetInstance() *gorm.DB {
	return db
}

// func (d *Database) GetInstance() *gorm.DB {
// 	return d.Instance
// }