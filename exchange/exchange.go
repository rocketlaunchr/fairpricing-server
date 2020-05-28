package exchangerate

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/rocketlaunchr/fairpricing/models"
)

// From https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html
const (
	xmlURL90 = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml" // Last 90 days only
	csvURL   = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist.zip"
)

// ErrNoCurrencyData indicates that no currency data exists for the sought after country code.
var ErrNoCurrencyData = errors.New("no currency data")

type rate struct {
	Value       *float64 // sometimes the csv data is N/A
	LastUpdated time.Time
}

var (
	lock         sync.RWMutex
	currentRates map[string]rate // Key is upper case currency code.  Val is 1 Euro = x
)

func init() {

	// Initial download
	lock.Lock()
	cr, err := UpdateExchangeRates()
	if err != nil {
		if currentRates == nil {
			panic(err)
		}

	}
	currentRates = cr
	lock.Unlock()

	// Update every 8 hours
	go func() {
		for range time.Tick(8 * time.Hour) {
			cr, err := UpdateExchangeRates()
			if err != nil {
				continue
			}
			lock.Lock()
			currentRates = cr
			lock.Unlock()
		}
	}()
}

func AllCurrencies() []string {
	curs := []string{}
	lock.RLock()
	for k := range currentRates {
		curs = append(curs, k)
	}
	lock.RUnlock()
	sort.Strings(curs)
	return curs
}

func GetExchangeRate(cur string) (float64, error) {
	lock.RLock()
	defer lock.RUnlock()

	r, exists := currentRates[cur]
	if !exists {
		return 0, ErrNoCurrencyData
	}

	if r.Value == nil {
		return 0, ErrNoCurrencyData
	}

	return *r.Value, nil
}

// ConvertExchangeRate converts price to toCurrency.
func ConvertExchangeRate(price models.Price, toCurrency string) (models.Price, error) {

	// Convert to Euro prices
	frR, err := GetExchangeRate(price.Currency)
	if err != nil {
		return models.Price{}, err
	}

	tEuro := price.Value / frR

	// Convert from Euro to target currency
	toR, err := GetExchangeRate(toCurrency)
	if err != nil {
		return models.Price{}, err
	}

	tTarget := toR * tEuro

	return models.Price{tTarget, toCurrency}, nil
}

// UpdateExchangeRates will redownload the latest exchange rate data from the European Central Bank.
//
// See: https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html
func UpdateExchangeRates() (map[string]rate, error) {

	currentRates := map[string]rate{
		"EUR": rate{&[]float64{1}[0], time.Now()}, // Base currency
	}

	// Download the zip file
	resp, err := http.Get(csvURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	z, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}

	zReader, err := z.File[0].Open()
	if err != nil {
		return nil, err
	}
	defer zReader.Close()

	// Decode csv
	r := csv.NewReader(zReader)

	var currencies []string

	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if i == 0 {
			currencies = record
			i++
		} else if i == 1 {

			var dt time.Time
			for j, col := range record {
				if j == 0 {
					dt = parseTime(col)
					continue
				}
				if col == "" {
					continue
				} else if col == "N/A" {
					// Ancient currencies
					continue
				}

				val, _ := strconv.ParseFloat(col, 64)
				currentRates[currencies[j]] = rate{&val, dt}
			}
			i++
		} else {
			break
		}
	}

	/*
	   USD	US dollar
	   JPY	Japanese yen
	   BGN	Bulgarian lev
	   CZK	Czech koruna
	   DKK	Danish krone
	   GBP	Pound sterling
	   HUF	Hungarian forint
	   PLN	Polish zloty
	   RON	Romanian leu
	   SEK	Swedish krona
	   CHF	Swiss franc
	   ISK	Icelandic krona
	   NOK	Norwegian krone
	   HRK	Croatian kuna
	   RUB	Russian rouble
	   TRY	Turkish lira
	   AUD	Australian dollar
	   BRL	Brazilian real
	   CAD	Canadian dollar
	   CNY	Chinese yuan renminbi
	   HKD	Hong Kong dollar
	   IDR	Indonesian rupiah
	   ILS	Israeli shekel
	   INR	Indian rupee
	   KRW	South Korean won
	   MXN	Mexican peso
	   MYR	Malaysian ringgit
	   NZD	New Zealand dollar
	   PHP	Philippine peso
	   SGD	Singapore dollar
	   THB	Thai baht
	   ZAR	South African rand
	*/

	// Add Hardcoded values
	currentRates["NGN"] = hardcode(429.170, "2020-03-30")
	currentRates["LKR"] = hardcode(207.896, "2020-03-30")

	return currentRates, nil
}

// LoadHardcode is used to load custom currencies and their value.
func LoadHardcode(cur string, val float64, t string) {
	lock.Lock()
	defer lock.Unlock()

	currentRates[cur] = hardcode(val, t)
}

func hardcode(val float64, t string) rate {
	if val == 0 {
		return rate{nil, parseTime(t)}
	}
	return rate{&val, parseTime(t)}
}

func parseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}
