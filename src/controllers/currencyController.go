package controllers

import (
	"challeng-bravo/src/models"
	"challeng-bravo/src/repositories"
	"challeng-bravo/src/requests"
	"challeng-bravo/src/responses"
	"challeng-bravo/src/validations"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

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

	if err = validations.ValidateCurrency(&currency); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(currency)

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

	if err = validations.ValidateCurrency(&currency); err != nil {
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

	to := strings.ToUpper(strings.TrimSpace(r.URL.Query().Get("to")))
	from := strings.ToUpper(strings.TrimSpace(r.URL.Query().Get("from")))
	amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 10)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = validations.ValidateConversionCurrency(to, from, amount)
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

	var conversion responses.CurrencyConversionResponse

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

	BTCApiCoinbase, err := requests.APICoinbase("BTC")

	ETHApiCoinbase, err := requests.APICoinbase("ETH")

	BRLInUSD := 1 / responseApiHGBrasil.Results.Currencies.USD.Sell
	_, err = UpdateValueInUSD("BRL", BRLInUSD)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	EURInUSD := responseApiHGBrasil.Results.Currencies.EUR.Sell / responseApiHGBrasil.Results.Currencies.USD.Sell

	_, err = UpdateValueInUSD("EUR", EURInUSD)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	BTCInUSD, err := strconv.ParseFloat(BTCApiCoinbase.Data.Amount, 10)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	_, err = UpdateValueInUSD("BTC", BTCInUSD)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	ETHInUSD, err := strconv.ParseFloat(ETHApiCoinbase.Data.Amount, 10)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	_, err = UpdateValueInUSD("ETH", ETHInUSD)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	var currentsQuotes responses.CurrentQuoteResponse

	currentsQuotes.Message = "Updated Quotes!"
	currentsQuotes.BRL = fmt.Sprintf("%f USD", BRLInUSD)
	currentsQuotes.EUR = fmt.Sprintf("%f USD", EURInUSD)
	currentsQuotes.BTC = fmt.Sprintf("%f USD", BTCInUSD)
	currentsQuotes.ETH = fmt.Sprintf("%f USD", ETHInUSD)
	currentsQuotes.Source = "Exchange data provided by HGBrasil and cryptocurrency by Coinbase"

	responses.JSON(w, http.StatusOK, currentsQuotes)

}

func UpdateValueInUSD(code string, ValueInUSD float64) (models.Currency, error) {

	currency, err := repositories.GetByCode(code)
	if err != nil {

		return currency, err
	}
	currency.Updated_at = time.Now()
	currency.ValueInUSD = ValueInUSD

	err = repositories.Update(currency, currency.ID)
	if err != nil {

		return currency, err
	}

	return currency, err
}
