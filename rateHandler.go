package main

import (
	"strings"

	"github.com/gofiber/fiber"
	exchangerate "github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/models"
)

type rates map[string]*models.Price

// FetchRates godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /rates/{base}/{date} [get]
func FetchRates(c *fiber.Ctx) {

	var r = rates{}

	// base currency eg AUD
	base := strings.ToUpper(c.Params("base"))
	err := validateCurrency(base)
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
			errorMsg := JsonErrorResponse{
				Error: &ApiError{
					Status: 404, Title: err.Error(),
				},
			}
			c.Status(404).JSON(errorMsg)

			return
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
			errorMsg := JsonErrorResponse{
				Error: &ApiError{
					Status: 500, Title: err.Error(),
				},
			}
			c.Status(500).JSON(errorMsg)

			return
		}
		r[cur] = &rate
	}

	response := JsonResponse{Data: r}
	c.JSON(response)

}
