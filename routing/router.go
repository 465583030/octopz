package routing

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Routes []Route
}

func (router *Router) AddRoute(route Route) {
	router.Routes = append(router.Routes, route)
}

func (r Router) ToMux() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range r.Routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
