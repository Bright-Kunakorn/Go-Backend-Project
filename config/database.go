package config

import (
	"golang-crud-gin/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() *gorm.DB {
	dsn := "postgresql://root:secret@localhost:5433?sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return DB
}




