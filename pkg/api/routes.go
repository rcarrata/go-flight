package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Defines a single route, HTTP method and the pattern
// the function that will execute when the route is called
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

// Defines the type Route, that is an array of Route structs
type Routes []Route

// Returns a pointer to a mux.Router we can use as a handler
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// Iterate over the routes and attach them to the router instance
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router

}

var routes = Routes{
	Route{
		Name:       "Index",
		Method:     "GET",
		Pattern:    "/",
		HandleFunc: Index,
	},
	Route{
		Name:       "FlightIndex",
		Method:     "GET",
		Pattern:    "/flights",
		HandleFunc: FlightIndex,
	},
	Route{
		Name:       "FlightShow",
		Method:     "GET",
		Pattern:    "/flight/{flightId}",
		HandleFunc: FlightShow,
	},
}
