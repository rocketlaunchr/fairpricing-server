package main

import (
	"sort"
	"strings"

	"github.com/gofiber/fiber"

	exchangerate "github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/fair"
)

var cnames = map[string]string{
	"EUR": "Euro",
	"USD": "US dollar",
	"JPY": "Japanese yen",
	"BGN": "Bulgarian lev",
	"CZK": "Czech koruna",
	"DKK": "Danish krone",
	"GBP": "Pound sterling",
	"HUF": "Hungarian forint",
	"PLN": "Polish zloty",
	"RON": "Romanian leu",
	"SEK": "Swedish krona",
	"CHF": "Swiss franc",
	"ISK": "Icelandic krona",
	"NOK": "Norwegian krone",
	"HRK": "Croatian kuna",
	"RUB": "Russian rouble",
	"TRY": "Turkish lira",
	"AUD": "Australian dollar",
	"BRL": "Brazilian real",
	"CAD": "Canadian dollar",
	"CNY": "Chinese yuan renminbi",
	"HKD": "Hong Kong dollar",
	"IDR": "Indonesian rupiah",
	"ILS": "Israeli shekel",
	"INR": "Indian rupee",
	"KRW": "South Korean won",
	"MXN": "Mexican peso",
	"MYR": "Malaysian ringgit",
	"NZD": "New Zealand dollar",
	"PHP": "Philippine peso",
	"SGD": "Singapore dollar",
	"THB": "Thai baht",
	"ZAR": "South African rand",
}

type Currency struct {
	Code string `json:"code"` // Currency code
	Name string `json:"name"`
}

type Country struct {
	Code string // ISO 3166-1 alpha-2 codes
	Name string // Title Case
}

type CurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
	Countries  []Country  `json:"countries"`
}

// CurrenciesHandler godoc
// @Summary Returns all currencies and countries recognized.
// @Produce json
// @Success 200 {object} CurrenciesResponse
// @Router /currencies [get]
func CurrenciesHandler(c *fiber.Ctx) {

	ac := exchangerate.AllCurrencies()
	cs := []Currency{}

	for _, c := range ac {
		cs = append(cs, Currency{
			Code: c,
			Name: cnames[c],
		})
	}

	scs := []Country{} // supported countries

	// We want all countries with a key in Ccodes and also exists in PppGDPperCapita
	for countryCode, v := range fair.Ccodes {
		country := strings.ToUpper(v.Name)
		if _, exists := fair.PppGDPperCapita[country]; exists {
			scs = append(scs, Country{
				Code: countryCode,
				Name: strings.Title(v.Name),
			})
		}
	}
	sort.Slice(scs, func(i, j int) bool { return scs[i].Name < scs[j].Name })

	resp := CurrenciesResponse{
		Currencies: cs,
		Countries:  scs,
	}

	c.JSON(resp)
}
