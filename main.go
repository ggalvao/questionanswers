package main

import (
	"log"
	"net/http"
)

func main() {
	db.Init()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
