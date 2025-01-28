package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	var err error
	dsn := "postgres://postgres:database.accesslog@localhost:5432/reminder_system"
	if Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	log.Print("Database connected successfully")

}
