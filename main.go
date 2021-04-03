package main

import (
	"log"
	"net/http"

	"github.com/rcarrata/go-flight/pkg/api"
)

// Flights -> Name, flight id, duration, Destination

func main() {

	port := ":8080"

	// Start the Router with the func NewRouter() in routes.go
	router := NewRouter()

	log.Printf("Server running in port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
