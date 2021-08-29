package controllers

import (
	"challeng-bravo/src/cache"
	"challeng-bravo/src/models"
	"challeng-bravo/src/repositories"
	"challeng-bravo/src/requests"
	"challeng-bravo/src/responses"
	"challeng-bravo/src/validations"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/gorilla/mux"
)

type valueInUSDInCache struct {
	ValueInUSD string `json:"valueInUSD"`
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

	saveValueInUSdInCache(currency.Code, currency.ValueInUSD, 30)

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

	saveValueInUSdInCache(currency.Code, currency.ValueInUSD, 30)

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	currency, err := repositories.GetById(ID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if currency.Code == "BRL" || currency.Code == "USD" || currency.Code == "ETH" || currency.Code == "BTC" {
		err = errors.New("Currency cannot be excluded.")
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = repositories.Delete(ID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	cache.Delete(currency.Code)

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

	var (
		toCurrencyInUSD   float64
		fromCurrencyInUSD float64
		toCurrency        models.Currency
		fromCurrency      models.Currency
	)

	toCurrencyInCache, _ := cache.Recover(to)

	if len(toCurrencyInCache) > 0 {

		data, _ := json.Marshal(toCurrencyInCache)
		var result valueInUSDInCache
		err := json.Unmarshal(data, &result)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		toCurrencyInUSD, _ = strconv.ParseFloat(result.ValueInUSD, 10)

	} else {
		toCurrency, err := repositories.GetByCode(to)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		toCurrencyInUSD = toCurrency.ValueInUSD

		saveValueInUSdInCache(to, toCurrencyInUSD, 30)

	}

	fromCurrencyInCache, _ := cache.Recover(from)

	if len(fromCurrencyInCache) > 0 {
		data, _ := json.Marshal(fromCurrencyInCache)
		var result valueInUSDInCache
		err := json.Unmarshal(data, &result)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		fromCurrencyInUSD, _ = strconv.ParseFloat(result.ValueInUSD, 10)
	} else {
		fromCurrency, err := repositories.GetByCode(from)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		fromCurrencyInUSD = fromCurrency.ValueInUSD

		saveValueInUSdInCache(from, fromCurrencyInUSD, 30)

	}

	convertedAmount := amount * (fromCurrencyInUSD / toCurrencyInUSD)

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

	currentQuoteInCache, _ := cache.Recover("currentQuote")
	if len(currentQuoteInCache) > 0 {
		data, _ := json.Marshal(currentQuoteInCache)
		var result responses.CurrentQuoteResponse
		err := json.Unmarshal(data, &result)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, result)

	} else {

		chanelBTCApiCoinbase := make(chan requests.ApiCoinbase)
		chanelETHApiCoinbase := make(chan requests.ApiCoinbase)
		chanelApiHGBrasil := make(chan requests.ApiHGBrasil)

		go requests.APICoinbase(chanelBTCApiCoinbase, "BTC")
		go requests.APICoinbase(chanelETHApiCoinbase, "ETH")
		go requests.APIHGBrasil(chanelApiHGBrasil)

		var (
			responseApiHGBrasil    requests.ApiHGBrasil
			responseApiCoinbaseBTC requests.ApiCoinbase
			responseApiCoinbaseETH requests.ApiCoinbase
		)

		for i := 0; i < 3; i++ {
			select {
			case apiHGBrasil := <-chanelApiHGBrasil:

				if apiHGBrasil.Results.Currencies.USD.Sell == 0 {

					responses.Error(w, http.StatusInternalServerError, errors.New("Error when trying to update quotes"))
					return
				}

				responseApiHGBrasil = apiHGBrasil

			case apiCoinbaseBTC := <-chanelBTCApiCoinbase:

				if apiCoinbaseBTC.Data.AmountFloat == 0 {

					responses.Error(w, http.StatusInternalServerError, errors.New("Error when trying to update quotes"))
					return

				}

				responseApiCoinbaseBTC = apiCoinbaseBTC

			case apiCoinbaseETH := <-chanelETHApiCoinbase:

				if apiCoinbaseETH.Data.AmountFloat == 0 {
					responses.Error(w, http.StatusInternalServerError, errors.New("Error when trying to update quotes"))
					return
				}

				responseApiCoinbaseETH = apiCoinbaseETH

			}
		}

		BRLInUSD := 1 / responseApiHGBrasil.Results.Currencies.USD.Sell
		err := UpdateValueInUSD("BRL", BRLInUSD)

		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}

		EURInUSD := responseApiHGBrasil.Results.Currencies.EUR.Sell / responseApiHGBrasil.Results.Currencies.USD.Sell

		err = UpdateValueInUSD("EUR", EURInUSD)

		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}

		BTCInUSD := responseApiCoinbaseBTC.Data.AmountFloat

		err = UpdateValueInUSD("BTC", BTCInUSD)

		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}

		ETHInUSD := responseApiCoinbaseETH.Data.AmountFloat

		err = UpdateValueInUSD("ETH", ETHInUSD)

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

		cache.Save("currentQuote", currentsQuotes, 30)

		responses.JSON(w, http.StatusOK, currentsQuotes)
	}

}

func UpdateValueInUSD(code string, ValueInUSD float64) error {

	currency, err := repositories.GetByCode(code)
	if err != nil {

		return err
	}
	currency.Updated_at = time.Now()
	currency.ValueInUSD = ValueInUSD

	err = repositories.Update(currency, currency.ID)
	if err != nil {

		return err
	}

	return err
}

func saveValueInUSdInCache(key string, value float64, expiryTimeInSeconds time.Duration) {

	data := valueInUSDInCache{}
	data.ValueInUSD = fmt.Sprintf("%f", value)
	cache.Save(key, data, 30)

}
