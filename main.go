package main

import (
	"avito-test/server"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func hello() {
	fmt.Println("hello every 5 secs")
	fmt.Printf("time: %v", time.Now().Format(time.DateTime))
}

func runCron() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(func() { hello() })

	s.StartBlocking()
}

func main() {
	log.Println("Starting App...")

	server := server.InitHttpServer(server.InitDatabase())

	go func() {
		runCron()
	}()

	server.RunServer()
}
