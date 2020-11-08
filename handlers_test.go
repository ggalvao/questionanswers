package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAuthorHandler(t *testing.T) {
	db.Init()
	reqBody := []byte(`{
		"FirstName": "Gabriel",
		"LastName": "Galvão",
		"Email": "ggalvao@gmail.com"
	}`)
	req, err := http.NewRequest("POST", "/addauthor", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddAuthorHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Id":0,"Email":"ggalvao@gmail.com","FirstName":"Gabriel","LastName":"Galvão","Questions":null,"Answers":null}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
func TestAddQuestionHandler(t *testing.T) {
	reqBody := []byte(`{
		"AuthorId": 0,
		"QuestionTitle": "Test2 Question",
		"QuestionSummary": "Summary for question",
		"Body": "ipsum lorem etc etc\nipsum lorem etc etc\nipsum lorem etc etc\nipsum lorem etc etc\nipsum lorem etc etc\n"
	}`)
	req, err := http.NewRequest("POST", "/addquestion", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddQuestionHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Id":0,"Title":"Test2 Question","Summary":"Summary for question","Body":"ipsum lorem etc etc\nipsum lorem etc etc\nipsum lorem etc etc\nipsum lorem etc etc\nipsum lorem etc etc\n","Answer":null}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
