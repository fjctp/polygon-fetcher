package fetcher

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

type FinData struct {
	Ticker string                  `json:"ticker"`
	Data   []models.StockFinancial `json:"data"`
}

// Fetch finanical data from Polygon using its client and
// save the data in JSON format
func FetchFinData(ticker string, count int, timeframe string,
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
	log.Printf("Get data since %s\n", date.Format("2006-01-02"))
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
				return FinData{}, errors.New("Invalid ticker")
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

// Read data from a JSON file and output a populated FinData structure
func ReadFile(path string) (FinData, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return FinData{}, err
	}

	var d FinData
	err = json.Unmarshal(content, &d)
	if err != nil {
		return FinData{}, err
	}
	return d, nil
}

// Write FinData to a JSON file
func (d FinData) Write(out_dir string) error {
	// write data to a JSON file named by ticker
	bytes, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	fname := d.Ticker + ".json"
	p := filepath.Join(out_dir, fname)

	log.Printf("Write data to %s\n", p)
	err = ioutil.WriteFile(filepath.Join(out_dir, fname), bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
