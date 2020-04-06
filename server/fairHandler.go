package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rocketlaunchr/fairpricing/fair"
	"github.com/rocketlaunchr/fairpricing/models"

)

//var countryCodes []string = []string{
//
//}

type convLocalPrice struct{
	OldLocalPrice *models.LocalPrice `json:"old_local_price"`
	NewLocalPrice *models.LocalPrice `json:"new_local_price"`
}

func FairExchange(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// locPrice e.g 100AUD@AU
	locPrice := strings.Split(strings.ToUpper(ps.ByName("locPrice")), "@")

	price, loc := locPrice[0], locPrice[1]

	currency := price[len(price)-3:]
	err := validateCurrency(currency)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := JsonErrorResponse{Error: &ApiError{Status: 400, Title: err.Error()}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}

	amount, err := strconv.ParseFloat(price[:len(price)-3], 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := JsonErrorResponse{Error: &ApiError{Status: 400, Title: err.Error()}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}

	toCountryCode := strings.ToUpper(ps.ByName("countryCode"))
	err = validateCountryCode(toCountryCode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := JsonErrorResponse{Error: &ApiError{Status: 400, Title: err.Error()}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}

	localPrice := models.LocalPrice{Price: models.Price{Value: amount, Currency: currency}, CountryCode: loc}

	toCurrency := strings.ToUpper(ps.ByName("currency"))
	toCurrency = strings.Trim(toCurrency, "/")

	var toCur string
	if toCurrency != "" {
		err = validateCurrency(toCurrency)
		if err == nil {
			toCur = toCurrency
		}
	}
	if toCur == "" {
		toCur = currency
	}

	np, err := fair.FairPrice(localPrice, toCountryCode, toCur)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := JsonErrorResponse{Error: &ApiError{Status: 500, Title: err.Error()}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
		return
	}

	response := JsonResponse{Data: &convLocalPrice{OldLocalPrice: &localPrice, NewLocalPrice: &np}}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}

}
