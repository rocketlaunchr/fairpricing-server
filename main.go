package main

import (
	"fmt"
	"strings"

	"github.com/rocketlaunchr/fairpricing/exchange"
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

//
func convertExchangeRate(fromPrice Price, toCurrency string) (Price, error) {

	// Convert to Euro prices
	frR, err := exchangerate.GetExchangeRate(fromPrice.Currency)
	if err != nil {
		return Price{}, err
	}

	tEuro := fromPrice.Value / frR

	// Convert from Euro to target currency
	toR, err := exchangerate.GetExchangeRate(toCurrency)
	if err != nil {
		return Price{}, err
	}

	tTarget := toR * tEuro

	return Price{tTarget, toCurrency}, nil
}

func fairPrice(fromLPrice LocalPrice, toCountryCode, toCurrency string) (LocalPrice, error) {

	// Step 1: Convert to USD
	cUSD, err := convertExchangeRate(fromLPrice.Price, "USD")
	if err != nil {
		return LocalPrice{}, err
	}

	// Step 2: What percentage of income is the cost?
	country := strings.ToUpper(ccodes[strings.ToLower(fromLPrice.CountryCode)])
	income := pppGDPperCapita[country]
	frac := cUSD.Value / income

	// Step 3: Convert frac to value in toCountryCode
	country = strings.ToUpper(ccodes[strings.ToLower(toCountryCode)])
	income = pppGDPperCapita[country]
	suggestedAmt := frac * income

	// Step 4: Convert suggestedAmt to toCurrency
	tCurrency, err := convertExchangeRate(Price{Value: suggestedAmt, Currency: "USD"}, toCurrency)
	if err != nil {
		return LocalPrice{}, err
	}

	return LocalPrice{Price: tCurrency, CountryCode: toCountryCode}, nil
}

func main() {
	p := LocalPrice{Price{10000, "AUD"}, "AU"}

	np, _ := fairPrice(p, "NZ", "NZD")
	fmt.Println("fair price:      ", np)
	discount, _ := convertExchangeRate(np.Price, "AUD")
	fmt.Println("discounted price:", discount)

}
