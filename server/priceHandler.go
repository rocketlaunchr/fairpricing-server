package main

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	exchangerate "github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/models"
)

type convPrice struct {
	OldPrice *models.Price `json:"old_price"`
	NewPrice *models.Price `json:"new_price"`
}

func PriceConvert(c *fiber.Ctx) {

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// price e.g. 10AUD
	price := strings.ToUpper(c.Params("price"))

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

	toCurrency := strings.ToUpper(c.Params("toCurrency"))
	err = validateCurrency(toCurrency)
	if err != nil {
		c.Status(404).Send(err.Error())
		return
	}

	date := c.Params("date") // not implemented for now
	_ = date

	p := models.Price{Value: amount, Currency: currency}

	newPrice, err := exchangerate.ConvertExchangeRate(p, toCurrency)
	if err != nil {
		c.Status(500).Send(err.Error())
		return
	}

	response := JsonResponse{Data: &convPrice{OldPrice: &p, NewPrice: &newPrice}}
	c.JSON(response)
}
