package main

import (
	"io"
	"log"
	"net/http"
)

var API_LINK string = "https://pokeapi.co/api/v2/location-area"

func commandMap() (string, error) {
	res, err := http.Get(API_LINK)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
