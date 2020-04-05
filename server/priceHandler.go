package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	exchangerate "github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/models"
	"net/http"
	"strconv"
	"strings"
)

type convPrice struct{
	OldPrice *models.Price `json:"old_price"`
	NewPrice *models.Price `json:"new_price"`
}

func PriceConvert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// price e.g. 10AUD
	price := strings.ToUpper(ps.ByName("price"))

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

	toCurrency := strings.ToUpper(ps.ByName("toCurrency"))
	err = validateCurrency(toCurrency)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: err.Error()}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}

	date := ps.ByName("date") // not implemented for now
	_ = date

	p := models.Price{Value: amount, Currency: currency}

	newPrice, err := exchangerate.ConvertExchangeRate(p, toCurrency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := JsonErrorResponse{Error: &ApiError{Status: 500, Title: err.Error()}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}

	response := JsonResponse{Data: &convPrice{OldPrice: &p, NewPrice: &newPrice}}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
