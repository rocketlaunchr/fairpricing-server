package main

import (
	"errors"
	"fmt"
	"strings"
	)

var currencies []string = []string{
	"USD", "JPY", "BGN", "CZK",
	"DKK", "GBP", "HUF", "PLN",
	"RON", "SEK", "CHF", "ISK",
	"NOK", "HRK", "RUB", "TRY",
	"AUD", "BRL", "CAD", "CNY",
	"HKD", "IDR", "ILS", "INR",
	"KRW", "MXN", "MYR", "NZD",
	"PHP", "SGD", "THB", "ZAR",
}

func validateCurrency(cur string) error {

	if len(cur) != 3 {
		return errors.New("invalid currency format")
	}

	if !strings.Contains(strings.Join(currencies, " "), cur) {
		return errors.New("invalid currency code format")
	}

	return nil
}

func validateCountryCode(cc string) error {
	if len(cc) != 2 {
		fmt.Println("I got here", cc)
		return errors.New("invalid country code format")
	}
	// more validations to be added (like list of country codes to search through)

	return nil
}
