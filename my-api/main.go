package main

import (
	"fmt"
	"log"
	"net/http"

	"git-amr-4.devtools.intel.com/gerrit/p/otc_cloud_ux-dbaas_api.git/endpoint"
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
