package routes

import (
	"fmt"
	"net/http"
)

var routesCurrency = []Router{
	{
		URI:    "/currency",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("currency teste")
		},
	},
}
