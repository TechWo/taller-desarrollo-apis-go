package endpoint

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techWomenCommunity/taller-desarrollo-apis-go/my-api/model"
)

// List fetchs all the available movies.
func List(w http.ResponseWriter, r *http.Request) {

	//TODO: Uses DAL fuction to list all the movie.
	moviesRaw := "TO DO"

	movies := model.ListMoviesResponse{
		Movies: moviesRaw,
	}

	log.Println("[RESPONSE] OK: The moview were fetched.")
	writeResponse(w, http.StatusOK, movies)

}

// Create inserts a new movie.
func Create(w http.ResponseWriter, r *http.Request) {

	var body *model.Movie
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close() // nolint: errcheck

	//TODO: Validate body inputs.
	// 1) Not nulls for name, year and director.
	// 2) Check that year is a valid one (format).

	//TODO: Uses DAL fuction for insert movies.

	log.Println("[RESPONSE] Created: The movie " + body.Name + " has been created successfully.")
	writeResponse(w, http.StatusCreated, nil)
}

// Delete deletes the specified movie.
func Delete(w http.ResponseWriter, r *http.Request) {

	movieID := mux.Vars(r)["id"]

	//TODO: Uses DAL fuction to delete the movie.

	log.Println("[RESPONSE] Deleted: The movie with the id " + movieID + " has been deleted successfully.")
	writeResponse(w, http.StatusOK, nil)

}

func writeResponse(w http.ResponseWriter, code int, object interface{}) {
	data, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	fmt.Fprintf(w, string(data))
}
