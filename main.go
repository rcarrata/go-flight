package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Flights -> Name, flight id, duration, Destination

func main() {

	port := ":8080"

	router := mux.NewRouter()

	log.Printf("Server running in port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
