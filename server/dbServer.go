package server

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dbURI := "postgres://postgres:zxc13243546@localhost:5432/User_slug_db"
	//dsn := "host=localhost user=postgres password=zxc13243546 dbname=User_slug_db port=5432 sslmode=disable TimeZone='Europe/Moscow parseTime=True"
	// dsn := "host=localhost user=postgres password=zxc13243546 dbname=User_slug_db port=5432 sslmode=disable TimeZone='Europe/Moscow"
	dbHandler, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while initializing Database %v", err)
	}

	return dbHandler
}
