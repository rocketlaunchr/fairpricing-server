package exchangerate

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/rocketlaunchr/fairpricing/models"
)

// From https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html
const (
	url    = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
	csvURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist.zip"
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

	// Update every 24 hours
	go func() {
		for range time.Tick(24 * time.Hour) {
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

	CurrentRates := map[string]rate{}

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
					dt, _ = time.Parse("2006-01-02", col)
					continue
				}
				if col == "" || col == "N/A" {
					continue
				}

				val, _ := strconv.ParseFloat(col, 64)
				CurrentRates[currencies[j]] = rate{&val, dt}
			}
			i++
		} else {
			break
		}
	}

	CurrentRates["EUR"] = rate{&[]float64{1}[0], time.Now()}
	return CurrentRates, nil
}
