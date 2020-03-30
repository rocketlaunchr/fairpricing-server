package main

import (
	"fmt"

	"github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/fair"
	"github.com/rocketlaunchr/fairpricing/models"
)

func main() {
	p := models.LocalPrice{models.Price{1, "AUD"}, "AU"}

	np, _ := fair.FairPrice(p, "NZ", "EUR")
	fmt.Println("fair price:      ", np)
	discount, _ := exchangerate.ConvertExchangeRate(np.Price, "AUD")
	fmt.Println("discounted price:", discount)

}
