package server

import (
	"github.com/NozomiSugiyama/go-api-server-practice/server/handler"
	"net/http"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"GET",
		"/users",
		handler.UserList,
	},
	Route{
		"POST",
		"/users",
		handler.UserCreate,
	},
	Route{
		"GET",
		"/users/{id}",
		handler.UserShow,
	},
	Route{
		"PUT",
		"/users/{id}",
		handler.UserUpdate,
	},
	Route{
		"DELETE",
		"/users/{id}",
		handler.UserDelete,
	},
}
