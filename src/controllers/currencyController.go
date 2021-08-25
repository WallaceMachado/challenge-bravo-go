package controllers

import (
	"challeng-bravo/src/repositories"
	"challeng-bravo/src/responses"

	"net/http"
)

func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := repositories.ListAll()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, currencies)

}
