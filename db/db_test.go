package db

import (
	"reflect"
	"testing"
	"time"

	"github.com/ggalvao/questionanswers/model"
)

var question = model.Question{
	AuthorID:    0,
	LastUpdated: time.Now(),
	TimeAdded:   time.Now(),
	Answer:      &model.Answer{},
	Body:        "Test body",
	ID:          1212,
	Summary:     "Test summary",
	Title:       "Test title",
}

var author = &model.Author{
	Answers:   []*model.Answer{},
	Email:     "test@email.com",
	FirstName: "Test",
	LastName:  "Test Last Name",
	Questions: []*model.Question{},
}

var answer = &model.Answer{
	AuthorID:     0,
	CreationTime: time.Now(),
	LastUpdated:  time.Now(),
	QuestionID:   0,
	Body:         "Test body",
}

func TestInit(t *testing.T) {
	Db.Init()
	if Db.questions == nil || Db.answers == nil || Db.authors == nil {
		t.Error("Failed initializing database")
	}
}

func TestAddGet(t *testing.T) {
	question = Db.AddQuestion(question)
	author = Db.AddAuthor(*author)
	answer = Db.AddAnswer(answer)
	dbQuestion := Db.questions[question.ID]
	dbAuthor := Db.authors[author.ID]
	dbAnswer := Db.answers[answer.ID]

	if !reflect.DeepEqual(*dbQuestion, question) {
		t.Error("Question add failed")
	}

	if !reflect.DeepEqual(*dbAuthor, *author) {
		t.Error("Author add failed")
	}

	if !reflect.DeepEqual(*dbAnswer, *answer) {
		t.Error("Answer add failed")
	}

	if !reflect.DeepEqual(question, *Db.GetQuestion(question.ID)) {
		t.Error("Question get failed")
	}
	if !reflect.DeepEqual(author, Db.GetAuthor(author.ID)) {
		t.Error("Question get failed")
	}
	if !reflect.DeepEqual(answer, Db.GetAnswer(answer.ID)) {
		t.Error("Question get failed")
	}
}

func TestDelete(t *testing.T) {
	q := Db.AddQuestion(question)
	Db.DeleteQuestion(q.ID)
	if Db.GetQuestion(q.ID) != nil {
		t.Error("Question not properly deleted!")
	}
	a := Db.AddAuthor(*author)
	Db.DeleteAuthor(a.ID)
	if Db.GetQuestion(a.ID) != nil {
		t.Error("Author not properly deleted!")
	}

	ans := Db.AddAnswer(answer)
	Db.DeleteAnswer(ans.ID)
	if Db.GetQuestion(ans.ID) != nil {
		t.Error("Answer not properly deleted!")
	}
}

func TestUpdate(t *testing.T) {
	oldQuestion := question
	oldQuestion.LastUpdated = time.Now()
	updatedQuestion := Db.UpdateQuestion(&oldQuestion)

	if !reflect.DeepEqual(&oldQuestion, updatedQuestion) {
		t.Error("Update failed")
	}

	oldAuthor := author
	oldAuthor.Email = "testchange@email.com"
	updatedAuthor := Db.UpdateAuthor(oldAuthor)

	if !reflect.DeepEqual(oldAuthor, updatedAuthor) {
		t.Error("Update failed")
	}

	oldAnswer := answer
	oldAnswer.LastUpdated = time.Now()
	updatedAnswer := Db.UpdateAnswer(oldAnswer)

	if !reflect.DeepEqual(oldAnswer, updatedAnswer) {
		t.Error("Update failed")
	}
}

func TestGetAllQuestions(t *testing.T) {
	Db.Init()
	Db.AddQuestion(question)
	Db.AddQuestion(question)
	Db.AddQuestion(question)
	Db.AddQuestion(question)

	questions := Db.GetAllQuestions()

	if len(questions) != 4 {
		t.Error("Get all questions failed")
	}
}
