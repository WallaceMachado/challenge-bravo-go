package controllers

import "net/http"

func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando um usuário"))
}
