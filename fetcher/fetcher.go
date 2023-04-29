package fetcher

import (
	"errors"
	"log"
	"os"

	"github.com/fjctp/polygon-fetcher/tickerData"
	polygon "github.com/polygon-io/client-go/rest"
)

type Fetcher struct {
	polygon *polygon.Client
}

// Get an instance of Fetcher
func New() (Fetcher, error) {
	key, isFound := os.LookupEnv("POLYGON_API_KEY")
	if isFound && key != "" {
		return Fetcher{polygon: polygon.New(key)}, nil
	} else {
		return Fetcher{}, errors.New(
			"Environment variable, POLYGON_API_KEY, is empty")
	}
}

func (f Fetcher) Fetch(ticker string) (tickerData.TickerData, error) {

	// Fetch finanical data
	log.Println("Fetch finanical data")
	fdata, err := FetchFinData(f, ticker, 100, "Q")
	if err != nil {
		return tickerData.TickerData{}, err
	}

	// Fetch dividend data
	log.Println("Fetch dividend data")
	ddata, err := FetchDividendData(f, ticker, 1000)
	if err != nil {
		return tickerData.TickerData{}, err
	}

	// Return data
	cname := fdata[0].CompanyName
	return tickerData.New(ticker, cname, fdata, ddata), nil
}
