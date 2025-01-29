package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=root dbname=simplerest port=5432 sslmode=disable"

	var err error
	// Open a connection to the PostgreSQL database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Print success message if connected
	log.Println("Successfully connected to the database")
}

func GetDB() *gorm.DB {
	return db
}
