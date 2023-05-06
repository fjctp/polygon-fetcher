package fetcher

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

// Fetch finanical data from Polygon using its client and
// save the data in JSON format
func FetchFinData(f Fetcher, ticker string, count int,
	timeframe string) ([]models.StockFinancial, error) {

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
	log.Printf("Get data since %s\n", date.Format("2006-01-02"))
	params := models.ListStockFinancialsParams{}.
		WithTicker(ticker).
		WithTimeframe(tf).
		WithPeriodOfReportDate(models.GTE, models.Date(date)).
		WithLimit(100).
		WithOrder(models.Asc)

	// make request
	iter := f.polygon.VX.ListStockFinancials(context.Background(), params)

	// read the next record and keep the data in a slice
	var data []models.StockFinancial
	for ind := 0; ind < (count * mulipler); ind++ {
		if !iter.Next() {
			if ind == 0 {
				return data, errors.New("Invalid ticker")
			}
			// reach the end of the results
			log.Print("End of record")
			break
		}

		item := iter.Item()
		log.Printf("%d, %s -> %s\n", ind, item.StartDate, item.EndDate)

		data = append(data, item)
	}

	return data, iter.Err()
}
