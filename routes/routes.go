package routes

import (
	"net/http"

	"github.com/ggalvao/questionanswers/handlers"
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
		handlers.AddQuestionHandler,
	},
	Route{
		"New Author",
		"POST",
		"/newauthor",
		handlers.AddAuthorHandler,
	},
	Route{
		"New Answer",
		"POST",
		"/newanswer",
		handlers.AddAnswerHandler,
	},
	Route{
		"Get Question",
		"GET",
		"/getquestion/{id}",
		handlers.GetQuestionHandler,
	},
	Route{
		"Get Author",
		"GET",
		"/getauthor/{id}",
		handlers.GetAuthorHandler,
	},
	Route{
		"Get Answer",
		"GET",
		"/getanswer/{id}",
		handlers.GetAnswerHandler,
	},
	Route{
		"Delete Question",
		"DELETE",
		"/deletequestion/{id}",
		handlers.DeleteQuestionHandler,
	},
	Route{
		"Delete Author",
		"DELETE",
		"/deleteauthor/{id}",
		handlers.DeleteAuthorHandler,
	},
	Route{
		"Delete Answer",
		"DELETE",
		"/deleteanswer/{id}",
		handlers.DeleteAnswerHandler,
	},
	Route{
		"Update Question",
		"PATCH",
		"/updatequestion/{id}",
		handlers.UpdateQuestionHandler,
	},
	Route{
		"Update Author",
		"PATCH",
		"/updateauthor/{id}",
		handlers.UpdateAuthorHandler,
	},
	Route{
		"Update Answer",
		"PATCH",
		"/updateanswer/{id}",
		handlers.UpdateAnswerHandler,
	},
	Route{
		"Get All Questions",
		"GET",
		"/getallquestions",
		handlers.GetAllQuestions,
	},
}
