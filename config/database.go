package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/morelmiles/booking-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Config() {
	e := godotenv.Load()

	if e != nil {
		log.Println(e)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbLink := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=require TimeZone=Africa/Nairobi password=%s", dbHost, username, dbName, dbPort, password)
	DB, err = gorm.Open(postgres.Open(dbLink), &gorm.Config{})

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Connected to the database")
	}

	DB.Debug().AutoMigrate(&models.User{})
}
