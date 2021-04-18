package api

import (
	"net/http"
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
