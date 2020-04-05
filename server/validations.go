package main

import (
	"errors"
	"strings"
	)

func validateCurrency(cur string) error {

	if len(cur) != 3 {
		return errors.New("invalid currency format")
	}

	if !strings.Contains(strings.Join(currencies, " "), cur) {
		return errors.New("invalid currency code format")
	}

	return nil
}
