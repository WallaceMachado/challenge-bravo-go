package router

import (
	"challeng-bravo/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {

	r := mux.NewRouter()
	return routes.SetUp(r)
}
