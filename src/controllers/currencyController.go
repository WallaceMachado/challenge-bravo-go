package controllers

import (
	"challeng-bravo/src/models"
	"challeng-bravo/src/repositories"
	"challeng-bravo/src/requests"
	"challeng-bravo/src/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/gorilla/mux"
)

type CurrencyFromResponse struct {
	Amount           string    `json:"amount"`
	QuoteUSD         string    `json:"quoteUSD"`
	QuoteUSDUpdateAt time.Time `json:"quoteUSDUpdateAt"`
}

type CurrencyToResponse struct {
	ConvertedAmount  string    `json:"convertedAmount"`
	QuoteUSD         string    `json:"quoteUSD"`
	QuoteUSDUpdateAt time.Time `json:"quoteUSDUpdateAt"`
}

type CurrencyConversionResponse struct {
	ConvertedAmount float64
	CurrencyFrom    CurrencyFromResponse
	CurrencyTo      CurrencyToResponse
}

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

	params := mux.Vars(r)
	ID := params["id"]

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
	params := mux.Vars(r)
	ID := params["id"]

	err := repositories.Delete(ID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func ConversionOfCurrency(w http.ResponseWriter, r *http.Request) {

	to := strings.ToUpper(r.URL.Query().Get("to"))
	from := strings.ToUpper(r.URL.Query().Get("from"))
	amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 10)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	toCurrency, err := repositories.GetByCode(to)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	fromCurrency, err := repositories.GetByCode(from)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	convertedAmount := amount * (fromCurrency.ValueInUSD / toCurrency.ValueInUSD)

	var conversion CurrencyConversionResponse

	conversion.ConvertedAmount = convertedAmount

	conversion.CurrencyFrom.Amount = fmt.Sprintf("%f %s", amount, from)
	conversion.CurrencyFrom.QuoteUSD = fmt.Sprintf("1 %s is worth %f USD", from, fromCurrency.ValueInUSD)
	conversion.CurrencyFrom.QuoteUSDUpdateAt = fromCurrency.Updated_at

	conversion.CurrencyTo.ConvertedAmount = fmt.Sprintf("%f %s", convertedAmount, to)
	conversion.CurrencyTo.QuoteUSD = fmt.Sprintf("1 %s is worth %f USD", to, toCurrency.ValueInUSD)
	conversion.CurrencyTo.QuoteUSDUpdateAt = toCurrency.Updated_at

	responses.JSON(w, http.StatusOK, conversion)

}

func CurrentQuote(w http.ResponseWriter, r *http.Request) {

	responseApiHGBrasil, err := requests.APIHGBrasil()
	resposeApiCoinbase, err := requests.APICoinbase()

	fmt.Println(responseApiHGBrasil.Results.Currencies.EUR.Sell, err, resposeApiCoinbase)

	responses.JSON(w, http.StatusOK, resposeApiCoinbase)

}
