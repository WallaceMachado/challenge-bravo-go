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
	{
		URI:    "/currency/{id}",
		Metodo: http.MethodPut,
		Funcao: controllers.UpdateCurrency,
	},
	{
		URI:    "/currency/{id}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeleteCurrency,
	},

	{
		URI:    "/currency/conversion",
		Metodo: http.MethodGet,
		Funcao: controllers.ConversionOfCurrency,
	},
	{
		URI:    "/currency/currentQuote",
		Metodo: http.MethodGet,
		Funcao: controllers.CurrentQuote,
	},
}
