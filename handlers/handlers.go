package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/ggalvao/questionanswers/db"
	"github.com/ggalvao/questionanswers/model"
	"github.com/ggalvao/questionanswers/util"
	"github.com/gorilla/mux"
)

func AddQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest model.RequestInformation
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		util.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	if db.Db.GetAuthor(jsonRequest.AuthorID) == nil {
		util.JSONResponse(w, http.StatusNotFound, model.ResponseInformation{Message: "Author not found"})
		return
	}
	q := model.Question{Title: jsonRequest.QuestionTitle, Summary: jsonRequest.QuestionSummary, Body: jsonRequest.Body, AuthorID: jsonRequest.AuthorID}
	q.TimeAdded = time.Now()
	q = db.Db.AddQuestion(q)
	author := db.Db.GetAuthor(jsonRequest.AuthorID)
	author.Questions = append(author.Questions, &q)
	db.Db.UpdateAuthor(author)
	util.JSONResponse(w, http.StatusOK, q)
}

func AddAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var newAuthor *model.Author
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &newAuthor); err != nil {
		util.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	newAuthor = db.Db.AddAuthor(*newAuthor)
	util.JSONResponse(w, http.StatusOK, newAuthor)
}

func AddAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest model.RequestInformation
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		util.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	question := db.Db.GetQuestion(jsonRequest.QuestionID)
	if question.Answer != nil {
		err := errors.New("question already answered. update instead")
		util.JSONResponse(w, http.StatusBadRequest, err)
		return
	}
	author := db.Db.GetAuthor(jsonRequest.AuthorID)
	a := &model.Answer{Body: jsonRequest.Body, QuestionID: question.ID}
	a.CreationTime = time.Now()
	a.LastUpdated = time.Now()
	a = db.Db.AddAnswer(a)

	question.Answer = a
	db.Db.UpdateQuestion(question)
	author.Answers = append(author.Answers, a)
	db.Db.UpdateAuthor(author)
	util.JSONResponse(w, http.StatusOK, *a)
}

func GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	question := db.Db.GetQuestion(id)
	if question == nil {
		err = errors.New("Question does not exist!")
		util.JSONResponse(w, http.StatusNotFound, err)
		return
	}
	util.JSONResponse(w, http.StatusOK, *question)
}

func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	author := db.Db.GetAuthor(id)
	if author == nil {
		util.JSONResponse(w, http.StatusNotFound, "err")
		return
	}
	util.JSONResponse(w, http.StatusOK, *author)
}

func GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	answer := db.Db.GetAnswer(id)
	util.JSONResponse(w, http.StatusOK, *answer)
}

func DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.Db.DeleteQuestion(id)
	util.JSONResponse(w, http.StatusOK, model.ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Question %d deleted", id)})
}

func DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.Db.DeleteAuthor(id)
	util.JSONResponse(w, http.StatusOK, model.ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Author %d deleted", id)})
}

func DeleteAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db.Db.DeleteAnswer(id)
	util.JSONResponse(w, http.StatusOK, model.ResponseInformation{Status: http.StatusOK, Message: fmt.Sprintf("Answer %d deleted", id)})
}

func UpdateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var updatedAuthor *model.Author

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &updatedAuthor); err != nil {
		util.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	author := db.Db.GetAuthor(id)
	author.Email = updatedAuthor.Email
	author.FirstName = updatedAuthor.FirstName
	author.LastName = updatedAuthor.LastName
	updatedAuthor = db.Db.UpdateAuthor(author)

	util.JSONResponse(w, http.StatusOK, updatedAuthor)
}

func UpdateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest model.RequestInformation
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		util.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	q := db.Db.GetQuestion(id)
	q.Body = jsonRequest.Body
	q.Summary = jsonRequest.QuestionSummary
	q.Title = jsonRequest.QuestionTitle
	q.LastUpdated = time.Now()

	q = db.Db.UpdateQuestion(q)
	util.JSONResponse(w, http.StatusOK, q)
}

func UpdateAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var jsonRequest model.RequestInformation
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &jsonRequest); err != nil {
		util.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	answer := db.Db.GetAnswer(id)
	answer.AuthorID = jsonRequest.AuthorID
	answer.Body = jsonRequest.Body
	answer.LastUpdated = time.Now()
	answer = db.Db.UpdateAnswer(answer)
	util.JSONResponse(w, http.StatusOK, answer)
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	util.JSONResponse(w, http.StatusOK, db.Db.GetAllQuestions())
}
