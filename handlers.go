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
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	if db.GetAuthor(jsonRequest.AuthorId) == nil {
		JSONResponse(w, http.StatusNotFound, ResponseInformation{Message: "Author not found"})
		return
	}
	q := Question{Title: jsonRequest.QuestionTitle, Summary: jsonRequest.QuestionSummary, Body: jsonRequest.Body, authorId: jsonRequest.AuthorId}
	q.timeAdded = time.Now()
	q = db.AddQuestion(q)
	author := db.GetAuthor(jsonRequest.AuthorId)
	author.Questions = append(author.Questions, &q)
	db.UpdateAuthor(author)
	JSONResponse(w, http.StatusOK, q)
}

func AddAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var newAuthor *Author
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &newAuthor); err != nil {
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	newAuthor = db.AddAuthor(*newAuthor)
	JSONResponse(w, http.StatusOK, newAuthor)
}

func AddAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest RequestInformation
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	question := db.GetQuestion(jsonRequest.QuestionId)
	if question.Answer != nil {
		err := errors.New("question already answered. update instead")
		JSONResponse(w, http.StatusBadRequest, err)
		return
	}
	author := db.GetAuthor(jsonRequest.AuthorId)
	a := &Answer{Body: jsonRequest.Body, questionId: question.Id}
	a.creationTime = time.Now()
	a.lastUpdated = time.Now()
	a = db.AddAnswer(a)

	question.Answer = a
	db.UpdateQuestion(question)
	author.Answers = append(author.Answers, a)
	db.UpdateAuthor(author)
	JSONResponse(w, http.StatusOK, *a)
}

func GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	question := db.GetQuestion(id)
	if question == nil {
		err = errors.New("Question does not exist!")
		JSONResponse(w, http.StatusNotFound, err)
		return
	}
	JSONResponse(w, http.StatusOK, *question)
}

func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	author := db.GetAuthor(id)
	JSONResponse(w, http.StatusOK, *author)
}

func GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	answer := db.GetAnswer(id)
	JSONResponse(w, http.StatusOK, *answer)
}

func DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.DeleteQuestion(id)
	JSONResponse(w, http.StatusOK, ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Question %d deleted", id)})
}

func DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.DeleteAuthor(id)
	JSONResponse(w, http.StatusOK, ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Author %d deleted", id)})
}

func DeleteAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.DeleteAnswer(id)
	JSONResponse(w, http.StatusOK, ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Answer %d deleted", id)})
}

func UpdateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var updatedAuthor *Author

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &updatedAuthor); err != nil {
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	author := db.GetAuthor(id)
	author.Email = updatedAuthor.Email
	author.FirstName = updatedAuthor.FirstName
	author.LastName = updatedAuthor.LastName
	updatedAuthor = db.UpdateAuthor(author)

	JSONResponse(w, http.StatusOK, updatedAuthor)
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
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	q := db.GetQuestion(id)
	q.Body = jsonRequest.Body
	q.Summary = jsonRequest.QuestionSummary
	q.Title = jsonRequest.QuestionTitle
	q.lastUpdated = time.Now()

	q = db.UpdateQuestion(q)
	JSONResponse(w, http.StatusOK, q)
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
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	answer := db.GetAnswer(id)
	answer.authorId = jsonRequest.AuthorId
	answer.Body = jsonRequest.Body
	answer.lastUpdated = time.Now()
	answer = db.UpdateAnswer(answer)
	JSONResponse(w, http.StatusOK, answer)
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, http.StatusOK, db.GetAllQuestions())
}
