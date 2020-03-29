package main

import (
	"time"
)

type rate struct {
	value       float64
	lastUpdated time.Time
}

var now = time.Now()

// From https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html
// 1 Euro = x
var exchangeRates = map[string]rate{
	"USD": {1.0977, now},
	"JPY": {119.36, now},
	"BGN": {1.9558, now},
	"CZK": {27.299, now},
	"DKK": {7.4606, now},
	"GBP": {0.89743, now},
	"HUF": {355.65, now},
	"PLN": {4.5306, now},
	"RON": {4.8375, now},
	"SEK": {11.0158, now},
	"CHF": {1.0581, now},
	"ISK": {154, now},
	"NOK": {11.6558, now},
	"HRK": {7.614, now},
	"RUB": {86.3819, now},
	"TRY": {7.0935, now},
	"AUD": {1.8209, now},
	"BRL": {5.5905, now},
	"CAD": {1.5521, now},
	"CNY": {7.7894, now},
	"HKD": {8.5095, now},
	"IDR": {17716.88, now},
	"ILS": {3.9413, now},
	"INR": {82.8695, now},
	"KRW": {1346.31, now},
	"MXN": {25.8329, now},
	"MYR": {4.7619, now},
	"NZD": {1.8548, now},
	"PHP": {56.125, now},
	"SGD": {1.5762, now},
	"THB": {35.769, now},
}
