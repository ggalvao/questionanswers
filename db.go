package main

type Database struct {
	questions []map[string]Question
	authors   []map[string]Author
	answers   []map[string]Answer
}

var db Database

func (d *Database) AddQuestion(q Question) {
	db.questions = append(db.questions, q)
}

func (d *Database) AddAuthor(a Author) {
	db.authors = append(db.authors, a)
}

func (d *Database) AddAnswer(a Answer) {
	db.answers = append(db.answers, a)
}
