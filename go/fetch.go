package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// An alias type for a map from cryptocurrency abbreviations to their floating-point exchange rate (from USD).
type ExchangeRates map[string]float64

// The structure of the response expected from the Coinbase API.
// This is used to when marshalling the JSON response.
// Note that the field names are uppercase to provide public access, while the JSON fields are actually lowercase.
// The "encoding/json" package makes this case conversion automatically.
type CoinbaseResponse struct {
	Data struct {
		Currency string
		Rates    map[string]string
	}
}

// Given an array of binary data representing the JSON response from the Coinbase API (see CoinbaseResponse),
// marshals that data using JSON then parses floating-points exchange rates from it.
func parseExchangeJson(data []byte) (ExchangeRates, error) {
	resObj := CoinbaseResponse{}
	jsonErr := json.Unmarshal(data, &resObj)
	if jsonErr != nil {
		return ExchangeRates{}, jsonErr
	}

	rates := make(ExchangeRates, len(resObj.Data.Rates))
	for key, value := range resObj.Data.Rates {
		rate, floatErr := strconv.ParseFloat(value, 64)
		if floatErr != nil {
			return rates, nil
		}
		rates[key] = rate
	}

	return rates, nil
}

// Fetches data from an API expected to follow the Coinbase exchange rate format,
// parsing it and returning a map from currency names to exchange rates (see ExchangeRates).
func fetchExchangeRates(uri string) (ExchangeRates, error) {
	rates := ExchangeRates{}
	httpClient := http.Client{
		// Default 5 second request timeout.
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return rates, err
	}

	// This probably isn't required by the endpoint, but I like to keep up appearances.
	req.Header.Set("Accept", "application/json")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return rates, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return rates, readErr
	}

	return parseExchangeJson(body)
}
