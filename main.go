package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fjctp/polygon-fetcher/fetcher"
	"github.com/fjctp/polygon-fetcher/report"
	"github.com/fjctp/polygon-fetcher/utils"
	"github.com/polygon-io/client-go/rest/models"
)

// Output directory for html and json files
const html_dir = "html"
const json_dir = "json"

// Hard limit on Polygon API
const max_item_per_req = 100
const max_num_of_req = 5
const max_item = max_item_per_req * max_num_of_req

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
	var num_year int
	flag.StringVar(&ticker, "ticker", "AAPL",
		"Get data for ticker. Default AAPL")
	flag.IntVar(&num_year, "num_year", 50,
		"Get data for the last number of years. Default: 50")
	flag.StringVar(&timeframe, "timeframe", "A",
		"A: annually, Q: quarterly. Default: A")
	flag.StringVar(&out_dir, "out_dir", "output",
		"Output directory. Default: output")
	flag.Parse()

	// Create directories
	out_path, err := filepath.Abs(out_dir)
	utils.CheckError(err)

	json_path := filepath.Join(out_path, json_dir)
	utils.MakeDirIfNotExist(json_path)

	html_path := filepath.Join(out_path, html_dir)
	utils.MakeDirIfNotExist(html_path)

	// Get a fetcher
	f, err := get_fetcher()
	utils.CheckError(err)

	// Fetch data
	ticker = strings.ToUpper(ticker)
	log.Printf("Fetch data for %s\n", ticker)
	ptimeframe := models.TFAnnual
	if timeframe == "Q" {
		ptimeframe = models.TFQuarterly
	}
	d, err := f.Fetch(ticker, num_year, ptimeframe)
	utils.CheckError(err)

	// Save data in JSON format
	err = d.Write(json_path)
	utils.CheckError(err)

	// Generate a report
	log.Printf("Generate report for %s\n", ticker)
	err = report.New(d, html_path)
	utils.CheckError(err)
}
