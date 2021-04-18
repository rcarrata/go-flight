package api

import (
	"encoding/json"
	"fmt"
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
