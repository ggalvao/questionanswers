package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	/*router.HandleFunc("/", homeLink)
	router.HandleFunc("/add/", Add).Methods("POST")
	router.HandleFunc("/delete/{id}", Delete).Methods("DELETE")
	router.HandleFunc("/update/{id}", Update).Methods("PATCH")
	router.HandleFunc("/get/{id}", Get).Methods("GET")
	router.HandleFunc("/get/", Get).Methods("GET")
	*/
	log.Fatal(http.ListenAndServe(":8080", router))
}
