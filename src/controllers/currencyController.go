package controllers

import (
	"challeng-bravo/src/models"
	"challeng-bravo/src/repositories"
	"challeng-bravo/src/responses"
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := repositories.ListAll()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, currencies)

}

func CreateCurrency(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var currency models.Currency

	if err = json.Unmarshal(bodyRequest, &currency); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	currencyID, err := repositories.Create(currency)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, currencyID)

}

func UpdateCurrency(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	ID := parametros["id"]

	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var currency models.Currency

	if err = json.Unmarshal(bodyRequest, &currency); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repositories.Update(currency, ID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID := parametros["id"]

	err := repositories.Delete(ID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
