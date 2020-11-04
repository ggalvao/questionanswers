package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Author etc
type Author struct {
	Email     string
	Questions []*Question
	Answers   []*Answer
}

// Answer etc
type Answer struct {
	Body         string
	Authors      []*Author
	creationTime time.Time
	lastUpdated  time.Time
}

// Question etc
type Question struct {
	Title     string
	Summary   string
	Body      string
	Author    Author
	answer    *Answer
	timeAdded time.Time
}

type RequestInformation struct {
	Username        string
	Body            string
	QuestionTitle   string
	QuestionSummary string
}

func AddQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var newQuestion RequestInformation
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &newQuestion); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	q := Question{Title: newQuestion.QuestionTitle, Summary: newQuestion.QuestionSummary, Body: newQuestion.Body, Author: db.Authors[newQuestion.Username]}
	newQuestion.timeAdded = time.Now()
	db.AddQuestion(newQuestion)
	fmt.Println(db)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newQuestion); err != nil {
		panic(err)
	}
}

func AddAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var newAuthor Author
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &newAuthor); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	db.AddAuthor(newAuthor)
	fmt.Println(db)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newAuthor); err != nil {
		panic(err)
	}
}

func AddAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var newAnswer Answer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &newAnswer); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	db.AddAnswer(newAnswer)
	fmt.Println(db)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newAnswer); err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/newquestion", AddQuestionHandler).Methods("POST")
	router.HandleFunc("/newauthor", AddAuthorHandler).Methods("POST")
	router.HandleFunc("/newanswer", AddAnswerHandler).Methods("POST")

	/*router.HandleFunc("/add/", Add).Methods("POST")
	router.HandleFunc("/delete/{id}", Delete).Methods("DELETE")
	router.HandleFunc("/update/{id}", Update).Methods("PATCH")
	router.HandleFunc("/get/{id}", Get).Methods("GET")
	router.HandleFunc("/get/", Get).Methods("GET")
	*/
	log.Fatal(http.ListenAndServe(":8080", router))
}
