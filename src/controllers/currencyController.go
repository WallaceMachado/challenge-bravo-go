package controllers

import (
	"challeng-bravo/src/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	dados, err := repositories.ListAll()

	if err != nil {
		fmt.Println("erro")
	}

	JSON(w, http.StatusOK, dados)

}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}
