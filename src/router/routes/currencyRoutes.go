package routes

import (
	"challeng-bravo/src/router/controllers"
	"net/http"
)

var routesCurrency = []Router{
	{
		URI:    "/currency",
		Metodo: http.MethodPost,
		Funcao: controllers.GetAllCurrencies,
	},
}
