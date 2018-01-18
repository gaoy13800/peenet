package http

import (
	"net/http"
	"github.com/gorilla/mux"
)

/**
	http 路由
 */

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
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Login",
		"GET",
		"/login",
		Login,
	},
	Route{
		"Logout",
		"GET",
		"/logout",
		Logout,
	},
	Route{
		"SaveDeviceData",
		"GET",
		"/saveclient",
		SaveDeviceData,
	},
}