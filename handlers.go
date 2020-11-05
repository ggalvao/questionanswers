package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func AddQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest RequestInformation
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	q := Question{Title: jsonRequest.QuestionTitle, Summary: jsonRequest.QuestionSummary, Body: jsonRequest.Body, Author: db.GetAuthor(jsonRequest.AuthorId)}
	q.timeAdded = time.Now()
	q = db.AddQuestion(q)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(q); err != nil {
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

	newAuthor = db.AddAuthor(newAuthor)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newAuthor); err != nil {
		panic(err)
	}
}

func AddAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest RequestInformation
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	question := db.GetQuestion(jsonRequest.QuestionId)
	if question.Answer != nil {
		err := errors.New("question already answered. update instead")
		panic(err)
	}
	author := db.GetAuthor(jsonRequest.AuthorId)
	authors := []*Author{author}
	a := Answer{Body: jsonRequest.Body, Question: question, Authors: authors}
	a.creationTime = time.Now()
	a.lastUpdated = time.Now()
	a = db.AddAnswer(a)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}

func GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db.GetQuestion(id))
}

func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db.GetAuthor(id))
}

func GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db.GetAnswer(id))
}

func DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.DeleteQuestion(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Question %d deleted", id)})
}

func DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.DeleteAuthor(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Author %d deleted", id)})
}

func DeleteAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.DeleteAnswer(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Answer %d deleted", id)})
}

func UpdateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var updatedAuthor Author

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &updatedAuthor); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	updatedAuthor = db.UpdateAuthor(updatedAuthor, id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(updatedAuthor); err != nil {
		panic(err)
	}
}

func UpdateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest RequestInformation
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	q := db.GetQuestion(id)
	q.Body = jsonRequest.Body
	q.Summary = jsonRequest.QuestionSummary
	q.Title = jsonRequest.QuestionTitle
	q.lastUpdated = time.Now()

	q = db.UpdateQuestion(*q, id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(q); err != nil {
		panic(err)
	}
}

func UpdateAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest RequestInformation
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	author := db.GetAuthor(jsonRequest.AuthorId)
	answer := db.GetAnswer(id)
	answer.Authors = append(answer.Authors, author)
	answer.Body = jsonRequest.Body
	answer.lastUpdated = time.Now()
	answer = db.UpdateAnswer(*answer, id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(answer); err != nil {
		panic(err)
	}
}
