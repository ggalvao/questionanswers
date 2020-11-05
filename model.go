package main

import "time"

// Author etc
type Author struct {
	Id        int
	Email     string
	FirstName string
	LastName  string
	Questions []*Question
	Answers   []*Answer
}

// Answer etc
type Answer struct {
	Id           int
	Body         string
	Authors      []*Author
	Question     *Question
	creationTime time.Time
	lastUpdated  time.Time
}

// Question etc
type Question struct {
	Id          int
	Title       string
	Summary     string
	Body        string
	Author      *Author
	Answer      *Answer
	timeAdded   time.Time
	lastUpdated time.Time
}
