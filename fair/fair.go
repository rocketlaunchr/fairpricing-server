package fair

import (
	"strings"

	"github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/models"
)

func FairPrice(fromLPrice models.LocalPrice, toCountryCode, toCurrency string) (models.LocalPrice, error) {

	// Step 1: Convert to USD
	cUSD, err := exchangerate.ConvertExchangeRate(fromLPrice.Price, "USD")
	if err != nil {
		return models.LocalPrice{}, err
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
	tCurrency, err := exchangerate.ConvertExchangeRate(models.Price{Value: suggestedAmt, Currency: "USD"}, toCurrency)
	if err != nil {
		return models.LocalPrice{}, err
	}

	return models.LocalPrice{Price: tCurrency, CountryCode: toCountryCode}, nil
}
