package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/fjctp/polygon-fetcher/fetcher"
	"github.com/fjctp/polygon-fetcher/report"
	"github.com/fjctp/polygon-fetcher/utils"
	"github.com/polygon-io/client-go/rest/models"
)

const max_item_per_req = 100
const max_num_of_req = 5
const max_item = max_item_per_req * max_num_of_req

const timeframe = models.TFAnnual

func get_fetcher() (*fetcher.Fetcher, error) {
	key, isFound := os.LookupEnv("POLYGON_API_KEY")
	if isFound {
		return fetcher.New(key), nil
	} else {
		return nil, errors.New("Invalid API key")
	}
}

func main() {
	// Define parameters
	var ticker, out_dir string
	var write_data bool
	var num_year int
	flag.StringVar(&ticker, "ticker", "AAPL", "Get data for ticker")
	flag.IntVar(&num_year, "num_year", 5, "Get data for the last number of years")
	flag.BoolVar(&write_data, "write_data", false, "Write data to out_dir in JSON format")
	flag.StringVar(&out_dir, "out_dir", "data", "Output directory")
	flag.Parse()

	// Get a fetcher
	f, err := get_fetcher()
	utils.Check_error(err)

	// Fetch data
	log.Printf("Fetch data for %s\n", ticker)
	d, err := f.Fetch(ticker, num_year, models.TFAnnual)
	utils.Check_error(err)

	// Save data in JSON format
	if write_data {
		err = d.Write(out_dir)
		utils.Check_error(err)
	}

	// Generate a report
	log.Printf("Generate report for %s\n", ticker)
	err = report.New(d)
	utils.Check_error(err)
}
