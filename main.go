package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestInformation struct {
	AuthorId        int
	QuestionId      int
	Body            string
	QuestionTitle   string
	QuestionSummary string
}
type ResponseInformation struct {
	Status  int
	Message string
}

func main() {
	db.Init()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/newquestion", AddQuestionHandler).Methods("POST")
	router.HandleFunc("/newauthor", AddAuthorHandler).Methods("POST")
	router.HandleFunc("/newanswer", AddAnswerHandler).Methods("POST")
	router.HandleFunc("/getquestion/{id}", GetQuestionHandler).Methods("GET")
	router.HandleFunc("/getauthor/{id}", GetAuthorHandler).Methods("GET")
	router.HandleFunc("/getanswer/{id}", GetAnswerHandler).Methods("GET")
	router.HandleFunc("/deletequestion/{id}", DeleteQuestionHandler).Methods("DELETE")
	router.HandleFunc("/deleteauthor/{id}", DeleteAuthorHandler).Methods("DELETE")
	router.HandleFunc("/deleteanswer/{id}", DeleteAnswerHandler).Methods("DELETE")
	router.HandleFunc("/updateauthor/{id}", UpdateAuthorHandler).Methods("PATCH")
	router.HandleFunc("/updatequestion/{id}", UpdateQuestionHandler).Methods("PATCH")
	router.HandleFunc("/updateanswer/{id}", UpdateAnswerHandler).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8080", router))
}
