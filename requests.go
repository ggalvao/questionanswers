package main

type RequestInformation struct {
	AuthorId        int
	QuestionId      int
	Body            string
	QuestionTitle   string
	QuestionSummary string
}
type ResponseInformation struct {
	Status  int
	Message string
}
