package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"New Question",
		"POST",
		"/newquestion",
		AddQuestionHandler,
	},
	Route{
		"New Author",
		"POST",
		"/newauthor",
		AddAuthorHandler,
	},
	Route{
		"New Answer",
		"POST",
		"/newanswer",
		AddAnswerHandler,
	},
	Route{
		"Get Question",
		"GET",
		"/getquestion/{id}",
		GetQuestionHandler,
	},
	Route{
		"Get Author",
		"GET",
		"/getauthor/{id}",
		GetAuthorHandler,
	},
	Route{
		"Get Answer",
		"GET",
		"/getanswer/{id}",
		GetAnswerHandler,
	},
	Route{
		"Delete Question",
		"DELETE",
		"/deletequestion/{id}",
		DeleteQuestionHandler,
	},
	Route{
		"Delete Author",
		"DELETE",
		"/deleteauthor/{id}",
		DeleteAuthorHandler,
	},
	Route{
		"Delete Answer",
		"DELETE",
		"/deleteanswer/{id}",
		DeleteAnswerHandler,
	},
	Route{
		"Update Question",
		"PATCH",
		"/updatequestion/{id}",
		UpdateQuestionHandler,
	},
	Route{
		"Update Author",
		"PATCH",
		"/updateauthor/{id}",
		UpdateAuthorHandler,
	},
	Route{
		"Update Answer",
		"PATCH",
		"/updateanswer/{id}",
		UpdateAnswerHandler,
	},
}
