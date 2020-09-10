package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const (
	helpString                = "Usage: fiatconv.exe amount from_currency to_currency\nExample: fiatconv 123.45 USD RUB"
	reqestFormat              = "https://api.exchangeratesapi.io/latest"
	errConnStatusMessage      = "Connection status error"
	errParsingJSONMessage     = "Error parsing JSON"
	errParsingCurrency        = "Currency %s not available\n"
	errParsingAmountMessage   = "Error parsing amount"
	resultMessage             = "According to European Central Bank %6.4f %s = %6.4f %s for %s\n"
	reachManMessage           = "%6.4f %v! Wow, you are rich man!\n"
	looserManMessage          = "Only %6.4f %v?! Loser!\n"
	availableCurruciesMessage = "Available currencies:\n"
)

func printAvailabeCurrency(currencies *map[string]interface{}) {
	fmt.Print(availableCurruciesMessage)
	for currencyName, _ := range *currencies {
		fmt.Fprintln(os.Stdout, currencyName)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Fprintln(os.Stdout, helpString)
		return
	}

	responce, err := http.Get(reqestFormat)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}
	if responce.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stdout, errConnStatusMessage, responce.StatusCode)
		return
	}

	var data interface{}
	if err := json.NewDecoder(responce.Body).Decode(&data); err != nil {
		fmt.Fprintln(os.Stdout, errParsingJSONMessage, err)
		return
	}

	rootData := data.(map[string]interface{})
	publishDate := rootData["date"].(string)
	currencyData := rootData["rates"].(map[string]interface{})

	fromValue := currencyData[args[1]]
	if fromValue == nil {
		fmt.Fprintf(os.Stdout, errParsingCurrency, args[1])
		printAvailabeCurrency(&currencyData)
		return
	}

	toValue := currencyData[args[2]]
	if toValue == nil {
		fmt.Fprintf(os.Stdout, errParsingCurrency, args[2])
		printAvailabeCurrency(&currencyData)
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprintln(os.Stdout, errParsingAmountMessage, args[0])
		return
	}

	convertionResult := amount / fromValue.(float64) * toValue.(float64)

	if amount > 100 {
		fmt.Fprintf(os.Stdout, reachManMessage, amount, args[1])
	}

	if amount < 10 {
		fmt.Fprintf(os.Stdout, looserManMessage, amount, args[1])
	}

	fmt.Fprintf(os.Stdout, resultMessage, amount, args[1], convertionResult, args[2], publishDate)
}
