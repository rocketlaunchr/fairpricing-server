package models

import (
	"fmt"
	"strings"
)

type Price struct {
	Value    float64 // smallest denomination (eg in cents for AUD)
	Currency string
}

func (p Price) String() string {
	return fmt.Sprintf("%.0f %s", p.Value, strings.ToUpper(p.Currency))
}

type LocalPrice struct {
	Price
	CountryCode string
}

func (p LocalPrice) String() string {
	return fmt.Sprintf("%.0f %s @ %s", p.Value, strings.ToUpper(p.Currency), strings.ToUpper(p.CountryCode))
}
