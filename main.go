package main

import (
	"avito-test/server"
	"log"
)

func main() {
	log.Println("Starting App...")
	log.Println("initializing config...")
	log.Println("Initializing database")
	log.Println("Initializing http server")

	server := server.InitHttpServer(server.InitDatabase())

	server.RunServer()
}
