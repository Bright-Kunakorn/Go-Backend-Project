package database

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func SetupDatabase(ctx context.Context) *gorm.DB {

	dsn := "postgresql://root:secret@localhost:5433?sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: false,
	})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	return DB
}
