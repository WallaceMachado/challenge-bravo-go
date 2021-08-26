package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiCoinbase struct {
	Data ApiCoinbaseCrypto `json:"data"`
}

type ApiCoinbaseCrypto struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func APICoinbase() (ApiCoinbase, error) {
	responseApiCoinbase := ApiCoinbase{}

	url := "https://api.coinbase.com/v2/prices/BTC-USD/sell"

	resp, err := http.Get(url)
	if err != nil {
		return responseApiCoinbase, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseApiCoinbase, err
	}
	if err != nil {

		return responseApiCoinbase, err

	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &responseApiCoinbase)
	if err != nil {
		return responseApiCoinbase, err
	}

	fmt.Println(responseApiCoinbase.Data.Amount)

	return responseApiCoinbase, nil
}
