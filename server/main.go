package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}

type ApiError struct {
	Status int16  `json:"status"`
	Title  string `json:"title"`
}

func main() {
	router := httprouter.New()
	// GET /rates/{base}/{date}?currencies=x,y,z
	router.GET("/rates/:base/:date", FetchRates)
	// GET /convert/{price}/{to currency}/{date}
	router.GET("/convert/:price/:toCurrency/:date", PriceConvert)
	// GET /fair/{price+location}/{country code}/{currency}
	router.GET("/fair/:locPrice/:countryCode/*currency", FairExchange)

	log.Fatal(http.ListenAndServe(":4321", router))
}
