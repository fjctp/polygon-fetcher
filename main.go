package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/fjctp/polygon-fetcher/fetcher"
	"github.com/fjctp/polygon-fetcher/report"
	"github.com/fjctp/polygon-fetcher/utils"
	"github.com/polygon-io/client-go/rest/models"
)

const max_item_per_req = 100
const max_num_of_req = 5
const max_item = max_item_per_req * max_num_of_req

const timeframe = models.TFQuarterly

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
	var ticker, timeframe, out_dir string
	var write_data bool
	var num_year int
	flag.StringVar(&ticker, "ticker", "AAPL",
		"Get data for ticker. Default AAPL")
	flag.IntVar(&num_year, "num_year", 50,
		"Get data for the last number of years. Default: 50")
	flag.StringVar(&timeframe, "timeframe", "A",
		"A: annually, Q: quarterly. Default: A")
	flag.BoolVar(&write_data, "write_data", true,
		"Write data to out_dir in JSON format. Default: true")
	flag.StringVar(&out_dir, "out_dir", "data",
		"Output directory. Default: data")
	flag.Parse()

	// Get a fetcher
	f, err := get_fetcher()
	utils.Check_error(err)

	// Fetch data
	ticker = strings.ToUpper(ticker)
	log.Printf("Fetch data for %s\n", ticker)
	ptimeframe := models.TFAnnual
	if timeframe == "Q" {
		ptimeframe = models.TFQuarterly
	}
	d, err := f.Fetch(ticker, num_year, ptimeframe)
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
