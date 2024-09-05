package main

import (
	"log"
)

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}
	log.Fatal(server.Listen(":8080"))

}
