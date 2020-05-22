package main

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/rocketlaunchr/fairpricing/fair"
	"github.com/rocketlaunchr/fairpricing/models"
)

type convLocalPrice struct {
	OldLocalPrice *models.LocalPrice `json:"old_local_price"`
	NewLocalPrice *models.LocalPrice `json:"new_local_price"`
}

// FairExchange godoc
// @Summary Show price in a local price of a given currency
// @Accept  json
// @Produce  json
// @Param locPrice path string true "Current Local Price"
// @Param countryCode path string true "Country Code to convert to"
// @Param currency path string true "Currency to convert to"
// @Success 200 {object} JsonResponse
// @Failure 400 {object} JsonErrorResponse
// @Failure 404 {object} JsonErrorResponse
// @Failure 500 {object} JsonErrorResponse
// @Router /fair/{locPrice}/{countryCode}/{currency} [get]
func FairExchange(c *fiber.Ctx) {

	// locPrice e.g 100AUD@AU
	locPrice := strings.Split(strings.ToUpper(c.Params("locPrice")), "@")

	price, loc := locPrice[0], locPrice[1]

	err := validateCountryCode(loc)
	if err != nil {
		errorMsg := JsonErrorResponse{
			Error: &ApiError{
				Status: 400, Title: err.Error(),
			},
		}
		c.Status(400).JSON(errorMsg)

		return
	}

	currency := price[len(price)-3:]
	err = validateCurrency(currency)
	if err != nil {
		errorMsg := JsonErrorResponse{
			Error: &ApiError{
				Status: 400, Title: err.Error(),
			},
		}
		c.Status(400).JSON(errorMsg)

		return
	}

	amount, err := strconv.ParseFloat(price[:len(price)-3], 64)
	if err != nil {
		errorMsg := JsonErrorResponse{
			Error: &ApiError{
				Status: 400, Title: err.Error(),
			},
		}
		c.Status(400).JSON(errorMsg)

		return
	}

	toCountryCode := strings.ToUpper(c.Params("countryCode"))
	err = validateCountryCode(toCountryCode)
	if err != nil {
		errorMsg := JsonErrorResponse{
			Error: &ApiError{
				Status: 400, Title: err.Error(),
			},
		}
		c.Status(400).JSON(errorMsg)

		return
	}

	localPrice := models.LocalPrice{Price: models.Price{Value: amount, Currency: currency}, CountryCode: loc}

	toCurrency := strings.ToUpper(c.Params("currency"))
	var toCur string
	if toCurrency != "" {
		err = validateCurrency(toCurrency)
		if err != nil {
			errorMsg := JsonErrorResponse{
				Error: &ApiError{
					Status: 404, Title: err.Error(),
				},
			}
			c.Status(404).JSON(errorMsg)
		}
		toCur = toCurrency

	} else {
		toCur = currency
	}

	np, err := fair.FairPrice(localPrice, toCountryCode, toCur)
	if err != nil {
		errorMsg := JsonErrorResponse{
			Error: &ApiError{
				Status: 500, Title: err.Error(),
			},
		}
		c.Status(500).JSON(errorMsg)

		return
	}

	response := JsonResponse{Data: &convLocalPrice{OldLocalPrice: &localPrice, NewLocalPrice: &np}}
	c.JSON(response)

}
