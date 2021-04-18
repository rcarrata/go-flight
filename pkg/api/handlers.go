package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

// Index Path API Rest
func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	p := path.Dir("./templates/")
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

//
func FlightIndex(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "These are the flights index")

	// Add Return Headers to the response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Add status OK to the response
	w.WriteHeader(http.StatusOK)

	// Call json.NewEncoder(w).Encoder(flights) to write JSON to the server
	err := json.NewEncoder(w).Encode(flights)
	if err != nil {
		panic(err)
	}
}

func FlightShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flightId := vars["flightId"]
	fmt.Fprintln(w, "Showing Flight num:", flightId)
}

// Create Flight func
func FlightCreate(w http.ResponseWriter, r *http.Request) {
	// create a new flight struct variable to use
	var flight Flight

	// Read all the body, and set a limitreader
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	// Close the content of the body
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Set all of body into an instance of a struct Flight
	if err := json.Unmarshal(body, &flight); err != nil {
		// If you receive an error send back a 422 Unprocessable Entity
		w.Header().Set("Content-Type", "application/json, charset=UTF-8")
		w.WriteHeader(422)

		// send back the err in a json string
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Create a flight and append into flights array
	f := dbCreateFlight(flight)
	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	// Sending back the representation of entity we created,
	// containing the id of the flight
	if err := json.NewEncoder(w).Encode(f); err != nil {
		panic(err)
	}
}
