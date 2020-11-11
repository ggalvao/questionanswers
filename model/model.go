package model

import "time"

// Author etc
type Author struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Questions []*Question
	Answers   []*Answer
}

// Answer etc
type Answer struct {
	ID           int
	Body         string
	AuthorID     int
	QuestionID   int
	CreationTime time.Time
	LastUpdated  time.Time
}

// Question etc
type Question struct {
	ID          int
	Title       string
	Summary     string
	Body        string
	AuthorID    int
	Answer      *Answer
	TimeAdded   time.Time
	LastUpdated time.Time
}
