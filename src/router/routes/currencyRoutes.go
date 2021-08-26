package routes

import (
	"challeng-bravo/src/controllers"
	"net/http"
)

var routesCurrency = []Router{
	{
		URI:    "/currency",
		Metodo: http.MethodGet,
		Funcao: controllers.GetAllCurrencies,
	},
	{
		URI:    "/currency",
		Metodo: http.MethodPost,
		Funcao: controllers.CreateCurrency,
	},
}
