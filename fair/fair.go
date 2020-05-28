package fair

import (
	"strings"

	"github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/models"
)

// FairPrice will convert p to the "fair price" for the country represented by toCountryCode.
// The process and rationale is defined in the README.
//
// If toCurrency is not provided, the currency of toCountryCode will be chosen by default.
//
// NOTE: This is not necessarily a direct currency converson.
func FairPrice(p models.LocalPrice, toCountryCode string, toCurrency ...string) (models.LocalPrice, error) {

	// Step 1: Convert to USD
	cUSD, err := exchangerate.ConvertExchangeRate(p.Price, "USD")
	if err != nil {
		return models.LocalPrice{}, err
	}

	// Step 2: What percentage of income is the cost?
	country := strings.ToUpper(Ccodes[strings.ToLower(p.CountryCode)].Name)
	income := PppGDPperCapita[country]
	frac := cUSD.Value / income

	// Step 3: Convert frac to value in toCountryCode
	country = strings.ToUpper(Ccodes[strings.ToLower(toCountryCode)].Name)
	income = PppGDPperCapita[country]
	suggestedAmt := frac * income

	// Step 4: Convert suggestedAmt to toCurrency
	if len(toCurrency) == 0 {
		toCurrency = append(toCurrency, Ccodes[strings.ToLower(toCountryCode)].Currency)
	}
	tCurrency, err := exchangerate.ConvertExchangeRate(models.Price{Value: suggestedAmt, Currency: "USD"}, toCurrency[0])
	if err != nil {
		return models.LocalPrice{}, err
	}

	return models.LocalPrice{Price: tCurrency, CountryCode: toCountryCode}, nil
}
