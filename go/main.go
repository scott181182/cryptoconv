/*
Cryptoconv prints USD converted to different cryptocurrencies.
When provided a USD amount and two cryptocurrencies (abbreviated), it will look up the exchange rates for those
currencies using the Coinbase API, then split the USD amount 70/30 and convert those amounts into the respective
currencies.

Usage:

	cryptoconv <USD amount> <currency 1> <currency 2>
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const COINBASE_URI = "https://api.coinbase.com/v2/exchange-rates?currency=USD"

func printUsage() {
	fmt.Println(`
USAGE
cryptoconv <USD amount> <currency 1> <currency 2>
	Converts 70% of USD amount to currency 1, and 30% to currency 2`)
}

func logConversion(usd float64, currency string, rates ExchangeRates) {
	// Being friendly and converting currency to uppercase.
	upperCurrency := strings.ToUpper(currency)

	rate, hasRate := rates[upperCurrency]
	if !hasRate {
		// We could fail here, but I thought it'd be nice to show the other conversion if it was successful.
		log.Printf("Could not find an exchange rate for '%s'. Please try with another.", currency)
	} else {
		// We use `strconv.FormatFloat` to get variable precision,
		// which makes it prettier or more accurate as necessary.
		log.Printf("$%.2f USD => %s %s", usd, strconv.FormatFloat(usd*rate, 'f', -1, 64), upperCurrency)
	}
}

func main() {
	// Remove timestamp from logger.
	log.SetFlags(0)

	// Get command-line arguments.
	args := os.Args
	if len(args) != 4 {
		// Account for the command name, which is the first argument.
		log.Printf("Expected 3 command line arguments, but found %d\n", len(args)-1)
		printUsage()
		return
	}

	usd, usdParseErr := strconv.ParseFloat(args[1], 64)
	if usdParseErr != nil {
		log.Println("There was an error parsing your USD amount. Please supply a number.")
		printUsage()
		return
	}
	if usd < 0 {
		log.Println("Nice try using a negative currency amount. I'll still convert it, but I'm onto you >.>")
	}

	currency1 := args[2]
	currency2 := args[3]

	rates, err := fetchExchangeRates(COINBASE_URI)
	if err != nil {
		log.Fatalf("There was an error fetching exchange rates:\n%+v\n", err)
	}

	logConversion(usd*0.7, currency1, rates)
	logConversion(usd*0.3, currency2, rates)
}
