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

func APICoinbase(canal chan<- ApiCoinbase, crypto string) {
	responseApiCoinbase := ApiCoinbase{}

	url := fmt.Sprintf("https://api.coinbase.com/v2/prices/%s-USD/sell", crypto)

	resp, err := http.Get(url)
	if err != nil {
		canal <- ApiCoinbase{}
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		canal <- ApiCoinbase{}
		return
	}
	if err != nil {

		canal <- ApiCoinbase{}
		return

	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &responseApiCoinbase)
	if err != nil {
		canal <- ApiCoinbase{}
		return
	}

	canal <- responseApiCoinbase

}
