package db

import "github.com/ggalvao/questionanswers/model"

type Database struct {
	questions                                  map[int]*model.Question
	authors                                    map[int]*model.Author
	answers                                    map[int]*model.Answer
	nextQuestionID, nextAuthorID, nextAnswerID int
}

var Db Database

func (d *Database) Init() {
	Db.answers = make(map[int]*model.Answer)
	Db.authors = make(map[int]*model.Author)
	Db.questions = make(map[int]*model.Question)
}
func (d *Database) AddQuestion(q model.Question) model.Question {
	q.ID = Db.nextQuestionID
	Db.questions[Db.nextQuestionID] = &q
	Db.nextQuestionID++
	return q
}

func (d *Database) GetQuestion(id int) *model.Question {
	question := Db.questions[id]
	return question
}

func (d *Database) DeleteQuestion(id int) {
	delete(Db.questions, id)
}

func (d *Database) UpdateQuestion(question *model.Question) *model.Question {
	Db.questions[question.ID] = question
	return question
}

func (d *Database) GetAllQuestions() []*model.Question {
	questionsSlice := []*model.Question{}
	for _, q := range Db.questions {
		questionsSlice = append(questionsSlice, q)
	}
	return questionsSlice
}
func (d *Database) AddAuthor(a model.Author) *model.Author {
	a.ID = Db.nextAuthorID
	Db.authors[Db.nextAuthorID] = &a
	Db.nextAuthorID++
	return &a
}

func (d *Database) GetAuthor(id int) *model.Author {
	author := Db.authors[id]
	return author
}

func (d *Database) DeleteAuthor(id int) {
	delete(Db.authors, id)
}

func (d *Database) UpdateAuthor(author *model.Author) *model.Author {
	Db.authors[author.ID] = author
	return author
}

func (d *Database) AddAnswer(a *model.Answer) *model.Answer {
	a.ID = Db.nextAnswerID
	Db.answers[Db.nextAnswerID] = a
	Db.nextAnswerID++
	return a
}

func (d *Database) GetAnswer(id int) *model.Answer {
	answer := Db.answers[id]
	return answer
}

func (d *Database) DeleteAnswer(id int) {
	delete(Db.answers, id)
}

func (d *Database) UpdateAnswer(answer *model.Answer) *model.Answer {
	Db.answers[answer.ID] = answer
	return answer
}
