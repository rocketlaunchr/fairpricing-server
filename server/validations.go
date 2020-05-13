package main

import (
	"fmt"
	"strings"

	"github.com/rocketlaunchr/fairpricing/fair"
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

func validateCurrency(cur string) error {

	if len(cur) != 3 {
		return fmt.Errorf("invalid currency code format: %s", cur)
	}

	for _, k := range currencies {
		if k == cur {
			return nil
		}
	}

	return fmt.Errorf("unavailable currency code format: %s", cur)
}

func validateCountryCode(cc string) error {
	if len(cc) != 2 {
		return fmt.Errorf("invalid country code format: %s", cc)
	}

	for k := range fair.Ccodes {
		if k == strings.ToLower(cc) {
			return nil
		}
	}

	return fmt.Errorf("unavailable country code format: %s", cc)
}
