package main

type Database struct {
	questions                                  map[int]*Question
	authors                                    map[int]*Author
	answers                                    map[int]*Answer
	nextQuestionID, nextAuthorID, nextAnswerID int
}

var db Database

func (d *Database) Init() {
	db.answers = make(map[int]*Answer)
	db.authors = make(map[int]*Author)
	db.questions = make(map[int]*Question)
}
func (d *Database) AddQuestion(q Question) Question {
	q.Id = db.nextQuestionID
	db.questions[db.nextQuestionID] = &q
	db.nextQuestionID++
	return q
}

func (d *Database) GetQuestion(id int) *Question {
	question := db.questions[id]
	return question
}

func (d *Database) DeleteQuestion(id int) {
	delete(db.questions, id)
}

func (d *Database) UpdateQuestion(question *Question) *Question {
	db.questions[question.Id] = question
	return question
}

func (d *Database) GetAllQuestions() []*Question {
	questionsSlice := []*Question{}
	for _, q := range db.questions {
		questionsSlice = append(questionsSlice, q)
	}
	return questionsSlice
}
func (d *Database) AddAuthor(a Author) *Author {
	a.Id = db.nextAuthorID
	db.authors[db.nextAuthorID] = &a
	db.nextAuthorID++
	return &a
}

func (d *Database) GetAuthor(id int) *Author {
	author := db.authors[id]
	return author
}

func (d *Database) DeleteAuthor(id int) {
	delete(db.authors, id)
}

func (d *Database) UpdateAuthor(author *Author) *Author {
	db.authors[author.Id] = author
	return author
}

func (d *Database) AddAnswer(a *Answer) *Answer {
	a.Id = db.nextAnswerID
	db.answers[db.nextAnswerID] = a
	db.nextAnswerID++
	return a
}

func (d *Database) GetAnswer(id int) *Answer {
	answer := db.answers[id]
	return answer
}

func (d *Database) DeleteAnswer(id int) {
	delete(db.answers, id)
}

func (d *Database) UpdateAnswer(answer *Answer) *Answer {
	db.answers[answer.Id] = answer
	return answer
}
