package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
