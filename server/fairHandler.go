package main

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/rocketlaunchr/fairpricing/fair"
	"github.com/rocketlaunchr/fairpricing/models"
)

//var countryCodes []string = []string{
//
//}

type convLocalPrice struct {
	OldLocalPrice *models.LocalPrice `json:"old_local_price"`
	NewLocalPrice *models.LocalPrice `json:"new_local_price"`
}

func FairExchange(c *fiber.Ctx) {

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// locPrice e.g 100AUD@AU
	locPrice := strings.Split(strings.ToUpper(c.Params("locPrice")), "@")

	price, loc := locPrice[0], locPrice[1]

	currency := price[len(price)-3:]
	err := validateCurrency(currency)
	if err != nil {
		c.Status(400).Send(err.Error())
		return
	}

	amount, err := strconv.ParseFloat(price[:len(price)-3], 64)
	if err != nil {
		c.Status(400).Send(err.Error())
		return
	}

	toCountryCode := strings.ToUpper(c.Params("countryCode"))
	err = validateCountryCode(toCountryCode)
	if err != nil {
		c.Status(400).Send(err.Error())
		return
	}

	localPrice := models.LocalPrice{Price: models.Price{Value: amount, Currency: currency}, CountryCode: loc}

	toCurrency := strings.ToUpper(c.Params("currency"))
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
		c.Status(500).Send(err.Error())
		return
	}

	response := JsonResponse{Data: &convLocalPrice{OldLocalPrice: &localPrice, NewLocalPrice: &np}}
	c.JSON(response)

}
