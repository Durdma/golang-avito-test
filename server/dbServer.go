package server

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	
	dbURI := "postgres://postgres:postgres@localhost:5432/User_slug_db"
	
	dbHandler, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while initializing Database %v", err)
	}

	return dbHandler
}
