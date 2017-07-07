package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/techWomenCommunity/taller-desarrollo-apis-go/my-api/endpoint"
	"github.com/techWomenCommunity/taller-desarrollo-apis-go/my-api/model"
)

const (
	port = "8080"
)

func main() {
	fmt.Println("Welcome to our first API!.")

	endpoint.MoviesDal = model.NewMovies()

	router := endpoint.Setup()
	http.Handle("/", router)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}

}
