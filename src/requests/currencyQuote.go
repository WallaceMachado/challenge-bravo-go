package requests

import (
	"challeng-bravo/src/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiHGBrasil struct {
	Results ApiHGBrasilCurrencies `json:"results"`
}

type ApiHGBrasilCurrencies struct {
	Currencies CurrenciesDefault `json:"currencies"`
}

type CurrenciesDefault struct {
	USD ApiHGBrasilCurrency `json:"USD"`
	EUR ApiHGBrasilCurrency `json:"EUR"`
}

type ApiHGBrasilCurrency struct {
	Buy       float64 `json:"Buy"`
	Name      string  `json:"Name"`
	Sell      float64 `json:"Sell"`
	Variation float64 `json:"Variation"`
}

func APIHGBrasil(canal chan<- ApiHGBrasil) {
	responseApiHGBrasil := ApiHGBrasil{}

	url := fmt.Sprintf("https://api.hgbrasil.com/finance/quotations?key=%s", config.KeyApiHGBRASIL)
	response, err := http.Get(url)

	if err != nil {
		canal <- ApiHGBrasil{}

		return

	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	err = json.Unmarshal(body, &responseApiHGBrasil)

	if err != nil {
		canal <- ApiHGBrasil{}

		return
	}

	canal <- responseApiHGBrasil
}
