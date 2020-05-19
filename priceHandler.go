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

// PriceConvert godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /convert/{price}/{toCurrency}/{date} [get]
func PriceConvert(c *fiber.Ctx) {

	// price e.g. 10AUD
	price := strings.ToUpper(c.Params("price"))

	currency := price[len(price)-3:]
	err := validateCurrency(currency)
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

	toCurrency := strings.ToUpper(c.Params("toCurrency"))
	err = validateCurrency(toCurrency)
	if err != nil {
		errorMsg := JsonErrorResponse{
			Error: &ApiError{
				Status: 404, Title: err.Error(),
			},
		}
		c.Status(404).JSON(errorMsg)

		return

	}

	date := c.Params("date") // not implemented for now
	_ = date

	p := models.Price{Value: amount, Currency: currency}

	newPrice, err := exchangerate.ConvertExchangeRate(p, toCurrency)
	if err != nil {
		errorMsg := JsonErrorResponse{
			Error: &ApiError{
				Status: 500, Title: err.Error(),
			},
		}
		c.Status(500).JSON(errorMsg)

		return
	}

	response := JsonResponse{Data: &convPrice{OldPrice: &p, NewPrice: &newPrice}}
	c.JSON(response)
}
