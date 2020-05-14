package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arsmn/fiber-swagger"
	"github.com/arsmn/fiber-swagger/example/docs"
	"github.com/gofiber/cors"
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
	app.Get("/rates/:base/:date?", FetchRates)
	// GET /convert/{price}/{to currency}/{date}
	app.Get("/convert/:price/:toCurrency/:date?", PriceConvert)
	// GET /fair/{price+location}/{country code}/{currency}
	app.Get("/fair/:locPrice/:countryCode/:currency?", FairExchange)
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	var (
		port = envInt("PORT", 4321)
		host = envString("HOST", fmt.Sprintf("http://localhost:%d", port))
	)

	app := fiber.New()
	docs.SwaggerInfo.Host = host

	app.Use(cors.New())
	app.Use("/docs", swagger.New(swagger.Config{ // custom
		// URL:         host + "/docs", //host + "/swagger.json",
		DeepLinking: true,
	}))
	setUpRoutes(app)
	app.Listen(port)
}

func envString(envVar string, defVal string) string {
	val, exists := os.LookupEnv(envVar)
	if !exists {
		return defVal
	}
	return val
}

func envInt(envVar string, defVal int) int {
	val, exists := os.LookupEnv(envVar)
	if !exists {
		return defVal
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return valInt
}
