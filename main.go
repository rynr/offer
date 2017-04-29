package main

import (
	"log"
	"net/http"
)

func main() {
	router := OffersRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
