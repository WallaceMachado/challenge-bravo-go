package routes

import (
	"challeng-bravo/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

func SetUp(r *mux.Router) *mux.Router {
	routes := routesCurrency

	for _, router := range routes {
		r.HandleFunc(router.URI, middleware.Loader(router.Funcao)).Methods(router.Metodo)
	}

	return r
}
