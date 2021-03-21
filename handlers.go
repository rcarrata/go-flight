package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

// Index Path API Rest
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//
func FlightIndex(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "These are the flights index")

	flights := Flights{
		Flight{
			Origin:      "Vancouver",
			Id:          1,
			Duration:    10,
			Destination: "Rome",
		},
		Flight{
			Origin:      "Madrid",
			Id:          2,
			Duration:    5,
			Destination: "Amsterdam",
		},
	}

	// Call json.NewEncoder(w).Encoder(flights) to write JSON to the server
	json.NewEncoder(w).Encode(flights)
}

func FlightShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flightId := vars["flightId"]
	fmt.Fprintln(w, "Showing Flight num:", flightId)
}
