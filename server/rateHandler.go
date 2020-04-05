package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	exchangerate "github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/models"
	"net/http"
	"strings"
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
}

var rates = make(map[string]*models.Price)

func FetchRates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// base currency eg AUD
	base := strings.ToUpper(ps.ByName("base"))
	err := validateCurrency(base)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: err.Error()}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}

	date := ps.ByName("date") // not implemented for now
	_ = date

	query := r.URL.Query()

	var (
		curs, curOpts []string
	)

	curncs, found := query["currencies"]
	if found && len(curncs) > 0 && curncs[0] != "" {
			curOpts = strings.Split(curncs[0], ",")

			for _, c := range curOpts {
				c = strings.ToUpper(c)

				err := validateCurrency(c)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
					response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: err.Error()}}
					if err := json.NewEncoder(w).Encode(response); err != nil {
						panic(err)
					}
					break
				}

				curs = append(curs, c)
			}
	} else {
		curs = currencies
	}

	for _, cur := range curs {
		//if cur == base {
		//	continue
		//} // skip adding the base currency to the currency rate list

		p := models.Price{Value: 1, Currency: cur}
		rate, err := exchangerate.ConvertExchangeRate(p, base)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := JsonErrorResponse{
				Error: &ApiError{
					Status: 500,
					Title: err.Error(),
				},
			}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				panic(err)
			}
			break
		}

		rates[cur] = &rate
	}

	response := JsonResponse{Data: rates}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
