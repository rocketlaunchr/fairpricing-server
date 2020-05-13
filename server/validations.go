package main

import (
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
	"EUR", "NGN", "LKR",
}

var countryCodes []string = []string{"", ""}

func validateCurrency(cur string) error {

	if len(cur) != 3 || !strings.Contains(strings.Join(currencies, " "), cur) {
		return fmt.Errorf("invalid currency code format: %s", cur)
	}

	return nil
}

func validateCountryCode(cc string) error {
	if len(cc) != 2 {
		return fmt.Errorf("invalid country code format: %s", cc)
	}
	// TODO: more validations to be added (like list of country codes to search through)

	return nil
}
