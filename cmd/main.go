package main

import (
	"github.com/Crud-application/cmd/server"
	"log"
)

func main() {

	s, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to start server, err: %v. Shutting down.", err)
	}
	// Set up routes
	s.SetupRoutes()

	// Run the server
	s.Run()
}
