package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/fjctp/polygon-fetcher/fetcher"
	"github.com/fjctp/polygon-fetcher/report"
	"github.com/fjctp/polygon-fetcher/utils"
)

// Output directory for html and json files
const html_dir = "html"
const json_dir = "json"

func main() {
	// Define parameters
	var num_year int
	var ticker, timeframe, output_dir string
	flag.StringVar(&ticker, "ticker", "AAPL",
		"Get data for ticker. Default AAPL")
	flag.IntVar(&num_year, "num_year", 50,
		"Get data for the last number of years. Default: 50")
	flag.StringVar(&timeframe, "timeframe", "A",
		"A: annually, Q: quarterly. Default: A")
	flag.StringVar(&output_dir, "output_dir", "output",
		"Output directory. Default: output")
	flag.Parse()

	// Validate inputs
	ticker = strings.ToUpper(ticker)
	timeframe = string([]rune(timeframe)[0])
	timeframe = strings.ToUpper(timeframe)
	if timeframe != "Q" {
		timeframe = "A"
	}

	// Create directories
	out_path, err := filepath.Abs(output_dir)
	utils.CheckError(err)

	json_path := filepath.Join(out_path, json_dir)
	utils.MakeDirIfNotExist(json_path)

	html_path := filepath.Join(out_path, html_dir)
	utils.MakeDirIfNotExist(html_path)

	// Fetch data
	log.Printf("Fetch data for %s\n", ticker)
	d, err := fetcher.FetchData(ticker, num_year, timeframe, json_path)
	utils.CheckError(err)

	// Generate a report
	log.Printf("Generate report for %s\n", ticker)
	err = report.New(ticker, d, html_path)
	utils.CheckError(err)
}
