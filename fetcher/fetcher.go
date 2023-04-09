package fetcher

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

// Get an instance of Fetcher
func NewFetcher() (*polygon.Client, error) {
	key, isFound := os.LookupEnv("POLYGON_API_KEY")
	if isFound {
		return polygon.New(key), nil
	} else {
		return nil, errors.New(
			"Environment variable, POLYGON_API_KEY, is empty")
	}
}

// Fetch finanical data from Polygon using its client and
// save the data in JSON format
func FetchData(ticker string, count int, timeframe string,
	output_dir string) (FinData, error) {
	f, err := NewFetcher()
	if err != nil {
		return FinData{}, err
	}

	var date time.Time
	var mulipler int
	var tf models.Timeframe
	now := time.Now()
	if timeframe == "Q" {
		// quarterly
		mulipler = 3
		date = now.AddDate(0, -count*mulipler, 0)
		tf = models.TFQuarterly
	} else {
		// anually
		mulipler = 1
		date = now.AddDate(-count, 0, 0)
		tf = models.TFAnnual
	}

	// set params
	params := models.ListStockFinancialsParams{}.
		WithTicker(ticker).
		WithTimeframe(tf).
		WithPeriodOfReportDate(models.GTE, models.Date(date)).
		WithLimit(100)

	// make request
	iter := f.VX.ListStockFinancials(context.Background(), params)

	// read the next record and keep the data in a slice
	var data []models.StockFinancial
	for ind := 0; ind < (count * mulipler); ind++ {
		if !iter.Next() {
			if ind == 0 {
				log.Fatal("Invalid ticker")
			}
			// reach the end of the results
			log.Print("End of record")
			break
		}

		item := iter.Item()
		log.Printf("%d, %s -> %s\n", ind, item.StartDate, item.EndDate)

		data = append(data, item)
	}

	// Save data in JSON format
	d := FinData{Ticker: ticker, Data: data}
	err = d.Write(output_dir)
	if err != nil {
		return d, err
	}

	return d, iter.Err()
}
