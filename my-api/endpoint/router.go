package endpoint

import "github.com/gorilla/mux"

// Setup enables the API endpoints and specified their handlers.
func Setup() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/movies", List).Methods("GET")
	router.HandleFunc("/movies", Create).Methods("POST")
	router.HandleFunc("/movies/{id}", Delete).Methods("DELETE")

	return router
}
