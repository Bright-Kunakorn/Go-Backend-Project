package config

import (
	"golang-crud-gin/helper"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := os.Getenv("DATABASE_CONNECTION_STRING")
	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	helper.ErrorPanic(err)

	return DB
}
