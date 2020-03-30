package models

import (
	"fmt"
	"strings"
)

// Price represents the price of an item.
type Price struct {
	// Value is always in the smallest denomination for a given currency.
	// Example: For AUD, in cents. Therefore 1 AUD has a value of 100.
	Value float64

	// Currency is the associated currency of the price.
	Currency string
}

// String implements fmt.Stringer
func (p Price) String() string {
	return fmt.Sprintf("%.0f %s", p.Value, strings.ToUpper(p.Currency))
}

// LocalPrice represents the price of an item in a particular country.
type LocalPrice struct {

	// Price represents the price of an item. The associated currency of the Price need not
	// be the official currency of the CountryCode.
	Price

	// Country code is the ISO 3166-1 alpha-2 codes.
	//
	// See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	CountryCode string
}

// String implements fmt.Stringer
func (p LocalPrice) String() string {
	return fmt.Sprintf("%.0f %s @ %s", p.Value, strings.ToUpper(p.Currency), strings.ToUpper(p.CountryCode))
}
