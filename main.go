package main

import (
	"log"
	"net/http"

	"github.com/ggalvao/questionanswers/db"
	"github.com/ggalvao/questionanswers/routes"
)

func main() {
	db.Db.Init()
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
