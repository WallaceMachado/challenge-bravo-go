package requests

import (
	"encoding/json"
	"fmt"
	"log"
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

// FazerRequisicaoComAutenticacao é utilizada para colocar o token na requisição
func APIHGBrasil() (ApiHGBrasil, error) {
	responseApiHGBrasil := ApiHGBrasil{}

	url := "https://api.hgbrasil.com/finance/quotations?key=b9524aa8"

	request, erro := http.NewRequest(http.MethodGet, url, nil)
	if erro != nil {
		fmt.Println(erro)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {

		return responseApiHGBrasil, err

	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseApiHGBrasil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseApiHGBrasil.Results.Currencies)

	return responseApiHGBrasil, nil
}
