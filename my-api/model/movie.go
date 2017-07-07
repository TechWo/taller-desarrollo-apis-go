package model

// Movie represents a movie structure.
type Movie struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// ListMoviesResponse represents the response body for the list movies request.
type ListMoviesResponse struct {
	Movies []Movie `json:"movies"`
}
