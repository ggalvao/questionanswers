package model

// RequestInformation describes a JSON request body
type RequestInformation struct {
	AuthorID        int
	QuestionID      int
	Body            string
	QuestionTitle   string
	QuestionSummary string
}

// ResponseInformation describes a JSON response
type ResponseInformation struct {
	Status  int
	Message string
}
