package main

import (
	"avito-test/server"
	"log"
)

func main() {
	log.Println("Starting App...")
	server := server.InitHttpServer()

	server.RunServer()
}
