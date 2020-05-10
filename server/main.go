package main

import (
	"github.com/gofiber/fiber"
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

func setUpRoutes(app *fiber.App) {
	// GET /rates/{base}/{date}?currencies=x,y,z
	app.Get("/rates/:base/:date", FetchRates)
	// GET /convert/{price}/{to currency}/{date}
	app.Get("/convert/:price/:toCurrency/:date", PriceConvert)
	// GET /fair/{price+location}/{country code}/{currency}
	app.Get("/fair/:locPrice/:countryCode/*currency", FairExchange)
}

func main() {
	app := fiber.New()

	setUpRoutes(app)
	app.Listen(4321)

}
