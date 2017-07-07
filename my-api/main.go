package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/techWomenCommunity/taller-desarrollo-apis-go/my-api/endpoint"
)

const (
	port = "8080"
)

func main() {
	fmt.Println("Welcome to our first API!.")

	router := endpoint.Setup()
	http.Handle("/", router)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}

}
