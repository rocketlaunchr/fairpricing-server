package main

import (
	"fmt"

	"github.com/rocketlaunchr/fairpricing/exchange"
	"github.com/rocketlaunchr/fairpricing/fair"
	"github.com/rocketlaunchr/fairpricing/models"
)

func main() {
	p := models.LocalPrice{models.Price{10000, "AUD"}, "AU"}

	np, _ := fair.FairPrice(p, "NZ", "NZD")
	fmt.Println("fair price:      ", np)
	discount, _ := exchangerate.ConvertExchangeRate(np.Price, "AUD")
	fmt.Println("discounted price:", discount)

}
