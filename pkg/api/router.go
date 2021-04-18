package api

import "github.com/gorilla/mux"

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
