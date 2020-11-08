package main

import (
	"reflect"
	"testing"
	"time"
)

var question = Question{
	authorId:    0,
	lastUpdated: time.Now(),
	timeAdded:   time.Now(),
	Answer:      &Answer{},
	Body:        "Test body",
	Id:          1212,
	Summary:     "Test summary",
	Title:       "Test title",
}

var author = &Author{
	Answers:   []*Answer{},
	Email:     "test@email.com",
	FirstName: "Test",
	LastName:  "Test Last Name",
	Questions: []*Question{},
}

var answer = &Answer{
	authorId:     0,
	creationTime: time.Now(),
	lastUpdated:  time.Now(),
	questionId:   0,
	Body:         "Test body",
}

func TestInit(t *testing.T) {
	db.Init()
	if db.questions == nil || db.answers == nil || db.authors == nil {
		t.Error("Failed initializing database")
	}
}

func TestAddGet(t *testing.T) {
	question = db.AddQuestion(question)
	author = db.AddAuthor(*author)
	answer = db.AddAnswer(answer)
	dbQuestion := db.questions[question.Id]
	dbAuthor := db.authors[author.Id]
	dbAnswer := db.answers[answer.Id]

	if !reflect.DeepEqual(*dbQuestion, question) {
		t.Error("Question add failed")
	}

	if !reflect.DeepEqual(*dbAuthor, *author) {
		t.Error("Author add failed")
	}

	if !reflect.DeepEqual(*dbAnswer, *answer) {
		t.Error("Answer add failed")
	}

	if !reflect.DeepEqual(question, *db.GetQuestion(question.Id)) {
		t.Error("Question get failed")
	}
	if !reflect.DeepEqual(author, db.GetAuthor(author.Id)) {
		t.Error("Question get failed")
	}
	if !reflect.DeepEqual(answer, db.GetAnswer(answer.Id)) {
		t.Error("Question get failed")
	}
}

func TestDelete(t *testing.T) {
	q := db.AddQuestion(question)
	db.DeleteQuestion(q.Id)
	if db.GetQuestion(q.Id) != nil {
		t.Error("Question not properly deleted!")
	}
	a := db.AddAuthor(*author)
	db.DeleteAuthor(a.Id)
	if db.GetQuestion(a.Id) != nil {
		t.Error("Author not properly deleted!")
	}

	ans := db.AddAnswer(answer)
	db.DeleteAnswer(ans.Id)
	if db.GetQuestion(ans.Id) != nil {
		t.Error("Answer not properly deleted!")
	}
}

func TestUpdate(t *testing.T) {
	oldQuestion := question
	oldQuestion.lastUpdated = time.Now()
	updatedQuestion := db.UpdateQuestion(&oldQuestion)

	if !reflect.DeepEqual(&oldQuestion, updatedQuestion) {
		t.Error("Update failed")
	}

	oldAuthor := author
	oldAuthor.Email = "testchange@email.com"
	updatedAuthor := db.UpdateAuthor(oldAuthor)

	if !reflect.DeepEqual(oldAuthor, updatedAuthor) {
		t.Error("Update failed")
	}

	oldAnswer := answer
	oldAnswer.lastUpdated = time.Now()
	updatedAnswer := db.UpdateAnswer(oldAnswer)

	if !reflect.DeepEqual(oldAnswer, updatedAnswer) {
		t.Error("Update failed")
	}
}

func TestGetAllQuestions(t *testing.T) {
	db.Init()
	db.AddQuestion(question)
	db.AddQuestion(question)
	db.AddQuestion(question)
	db.AddQuestion(question)

	questions := db.GetAllQuestions()

	if len(questions) != 4 {
		t.Error("Get all questions failed")
	}
}
