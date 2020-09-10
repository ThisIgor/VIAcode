package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const (
	helpString              = "Parameters: amount from_currency to_currency\nExample: fiatconv 123.45 USD RUB"
	reqestFormat            = "https://api.exchangeratesapi.io/latest"
	errConnStatusMessage    = "Connection status error"
	errParsingJSONMessage   = "Error parsing JSON"
	errPatsingCurrency      = "No value for"
	errParsingAmountMessage = "Error parsing amount"
	resultMessage           = "According to European Central Bank %6.4f %s = %6.4f %s for %s\n"
	reachManMessage         = "%6.4f %v! Wow, you are rich man!\n"
	looserManMessage        = "Only %6.4f %v?! Loser!\n"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println(helpString)
		return
	}

	responce, err := http.Get(reqestFormat)
	if err != nil {
		fmt.Println(err)
		return
	}
	if responce.StatusCode != http.StatusOK {
		fmt.Println(errConnStatusMessage, responce.StatusCode)
		return
	}

	var data interface{}
	if err := json.NewDecoder(responce.Body).Decode(&data); err != nil {
		fmt.Println(errParsingJSONMessage, err)
		return
	}

	mainData := data.(map[string]interface{})
	publishDate := mainData["date"].(string)
	currencyData := mainData["rates"].(map[string]interface{})

	fromValue := currencyData[args[1]]
	if fromValue == nil {
		fmt.Println(errPatsingCurrency, args[1])
		return
	}

	toValue := currencyData[args[2]]
	if toValue == nil {
		fmt.Println(errPatsingCurrency, args[2])
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println(errParsingAmountMessage, args[0])
		return
	}

	convertionResult := amount / fromValue.(float64) * toValue.(float64)

	if amount > 100 {
		fmt.Printf(reachManMessage, amount, args[1])
	}

	if amount < 10 {
		fmt.Printf(looserManMessage, amount, args[1])
	}

	fmt.Printf(resultMessage, amount, args[1], convertionResult, args[2], publishDate)
}
