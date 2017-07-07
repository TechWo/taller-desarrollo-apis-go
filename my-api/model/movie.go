package model

import (
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Movie represents a movie structure.
type Movie struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// ListMoviesResponse represents the response body for the list movies request.
type ListMoviesResponse struct {
	Movies []*Movie `json:"movies"`
}

//Movies moviesDAL
type Movies struct {
	moviesContainer map[string]*Movie
}

func NewMovies() *Movies {
	return &Movies{
		moviesContainer: make(map[string]*Movie),
	}
}

//List returns all the movie
func (t *Movies) List() []*Movie {

	movies := make([]*Movie, len(t.moviesContainer))
	i := 0
	for _, v := range t.moviesContainer {
		movies[i] = v
		i++
	}

	return movies
}

//Insert insert a movie
func (t *Movies) Insert(m *Movie) (string, error) {
	id, e := t.newUUID()

	if e != nil {
		return "", e
	}

	m.ID = id
	t.moviesContainer[id] = m

	return id, nil
}

//FindByID find a movie
func (t *Movies) FindByID(id string) *Movie {
	return t.moviesContainer[id]
}

//Delete deletes a movie
func (t *Movies) Delete(id string) {
	delete(t.moviesContainer, id)
}

//Search does lookup to the movie Container
func (t *Movies) Search(f string, v string) *Movie {

	if f == "Name" {
		return t.searchByName(v)
	} else if f == "Year" {
		i, _ := strconv.Atoi(v)
		return t.searchByYear(i)
	} else if f == "Director" {
		return t.searchByDirector(v)
	} else if f == "ID" {
		return t.FindByID(v)
	}

	return nil
}

//SearchByName searches a Movie byField, except Year
func (t *Movies) searchByName(value string) *Movie {
	for _, v := range t.moviesContainer {
		if strings.ContainsAny(v.Name, value) {
			return v
		}
	}
	return nil
}

//SearchByYear searches a Movie byField, except Year
func (t *Movies) searchByYear(value int) *Movie {
	for _, v := range t.moviesContainer {
		if v.Year == value {
			return v
		}
	}
	return nil
}

//SearchByDirector searches a Movie byField, except Year
func (t *Movies) searchByDirector(value string) *Movie {
	for _, v := range t.moviesContainer {
		if strings.ContainsAny(v.Director, value) {
			return v
		}
	}
	return nil
}

// newUUID generates a random UUID according to RFC 4122
func (t *Movies) newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
