package responses

import (
	"time"
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
	ConvertedAmount float64              `json:"convertedAmount"`
	CurrencyFrom    CurrencyFromResponse `json:"currencyFrom"`
	CurrencyTo      CurrencyToResponse   `json:"currencyTo"`
}

type CurrentQuoteResponse struct {
	Message string `json:"message"`
	BRL     string
	EUR     string
	BTC     string
	ETH     string
	Source  string `json:"source"`
}
