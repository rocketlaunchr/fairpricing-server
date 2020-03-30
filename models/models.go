package models

import (
	"fmt"
	"strings"
)

// Price represents the price of an item.
type Price struct {

	// Value is the numerical price of an item in a given currency.
	Value float64

	// Currency is the associated currency of the price.
	Currency string
}

// ValueISO4217 converts the Value to the correct number of decimal places.
//
// See: https://en.wikipedia.org/wiki/ISO_4217
func (p Price) ValueISO4217() string {
	switch p.Currency {
	case "BIF", "CLP", "DJF", "GNF", "ISK", "JPY", "KMF", "KRW", "PYG", "RWF", "UGX", "VND", "VUV", "XAF", "XOF", "XPF":
		return fmt.Sprintf("%.0f", p.Value)
	case "BHD", "IQD", "JOD", "KWD", "TND":
		return fmt.Sprintf("%.3f", p.Value)
	default:
		return fmt.Sprintf("%.2f", p.Value)
	}
}

// String implements fmt.Stringer
func (p Price) String() string {
	return fmt.Sprintf("%s %s", p.ValueISO4217(), strings.ToUpper(p.Currency))
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
	return fmt.Sprintf("%s %s @ %s", p.ValueISO4217(), strings.ToUpper(p.Currency), strings.ToUpper(p.CountryCode))
}
