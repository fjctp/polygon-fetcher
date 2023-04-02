package fetcher

import (
	"context"
	"log"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

// A wrapper of polygon.Client
type Fetcher struct {
	client polygon.Client
}

// Get an instance of Fetcher
func New(key string) *Fetcher {
	c := polygon.New(key)
	return &Fetcher{client: *c}
}

// Fetch finanical data from Polygon using its client
func (f Fetcher) Fetch(ticker string, count int, tf models.Timeframe) (FinData, error) {
	var date time.Time
	var mulipler int
	now := time.Now()
	if tf == models.TFAnnual {
		// anually
		mulipler = 1
		date = now.AddDate(-count, 0, 0)
	} else {
		// quarterly
		mulipler = 3
		date = now.AddDate(0, -count*mulipler, 0)
	}

	// set params
	params := models.ListStockFinancialsParams{}.
		WithTicker(ticker).
		WithTimeframe(tf).
		WithPeriodOfReportDate(models.GTE, models.Date(date))

	// make request
	iter := f.client.VX.ListStockFinancials(context.Background(), params)

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

	return FinData{Ticker: ticker, Data: data}, iter.Err()
}
