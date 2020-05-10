package main

import (
	"strings"

	"github.com/gofiber/fiber"
	exchangerate "github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/models"
)

var rates = make(map[string]*models.Price)

func FetchRates(c *fiber.Ctx) {

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// base currency eg AUD
	base := strings.ToUpper(c.Params("base"))
	err := validateCurrency(base)
	if err != nil {
		c.Status(404).Send(err.Error())

		return
	}

	date := c.Params("date") // not implemented for now
	_ = date

	var (
		curs, curOpts []string
	)

	// curncs, found := query["currencies"]
	curncs := c.Query("currencies")
	curOpts = strings.Split(curncs, ",")
	for _, cur := range curOpts {
		curr := strings.ToUpper(cur)

		err := validateCurrency(curr)
		if err != nil {
			c.Status(404).Send(err.Error())
			break
		}

		curs = append(curs, curr)
	}
	for _, cur := range curs {
		//if cur == base {
		//	continue
		//} // skip adding the base currency to the currency rate list

		p := models.Price{Value: 1, Currency: cur}
		rate, err := exchangerate.ConvertExchangeRate(p, base)
		if err != nil {
			c.Status(500).Send(err.Error())

			break
		}

		rates[cur] = &rate
	}

	response := JsonResponse{Data: rates}
	c.JSON(response)

}
