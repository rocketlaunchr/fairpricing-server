package main

import (
	"fmt"
	"strings"

	exchangerate "github.com/rocketlaunchr/fairpricing/exchange"
	fair "github.com/rocketlaunchr/fairpricing/fair"
)

func validateCurrency(cur string) error {

	if len(cur) != 3 {
		return fmt.Errorf("invalid currency code format: %s", cur)
	}

	for k := range exchangerate.CurrentRates {
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
