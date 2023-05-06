package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fjctp/polygon-fetcher/fetcher"
	"github.com/fjctp/polygon-fetcher/middleware"
	"github.com/fjctp/polygon-fetcher/report"
	"github.com/fjctp/polygon-fetcher/tickerData"
	"github.com/fjctp/polygon-fetcher/utils"
)

// Output directory for html and json files
const html_dir = "html"
const json_dir = "json"

func updateJsonReports(json_path string, html_path string) middleware.Updater {
	// Update JSON data or HTML report if it is older than a certain amount of days
	olderThanDays := 30 * 3

	return func(ticker string) error {
		// Validate inputs
		ticker = strings.ToUpper(ticker)

		// Fetch data
		f, err := fetcher.New()
		var d tickerData.TickerData
		t_json_path := filepath.Join(json_path, ticker+".json")
		update_json := !utils.Exist(t_json_path) ||
			utils.FileOlderThan(t_json_path, 0, 0, olderThanDays)
		if update_json {
			log.Printf("Data does not exist or is outdated for for %s, fetching...\n", ticker)
			d, err = f.Fetch(ticker)
			if err != nil {
				return err
			}
			d.Write(json_path)
		} else {
			log.Printf("Data exists for for %s\n", ticker)
			d, err = tickerData.ReadFile(t_json_path)
			if err != nil {
				return err
			}
		}

		// Generate a report
		t_html_path := filepath.Join(html_path, ticker+".html")
		update_html := !utils.Exist(t_html_path) ||
			utils.FileOlderThan(t_html_path, 0, 0, olderThanDays)
		if update_json || update_html {
			log.Printf("Report does not exist or is outdated for %s, generating...\n", ticker)
			err = report.New(d, html_path)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func main() {
	// Define parameters
	var output_dir string
	var port int
	flag.StringVar(&output_dir, "output_dir", "output",
		"Output directory. Default: output")
	flag.IntVar(&port, "port", 80,
		"Server listent port. Default: 80")
	flag.Parse()

	// Create directories
	out_path, err := filepath.Abs(output_dir)
	utils.CheckError(err)

	json_path := filepath.Join(out_path, json_dir)
	utils.CheckError(utils.MakeDirIfNotExist(json_path))

	html_path := filepath.Join(out_path, html_dir)
	utils.CheckError(utils.MakeDirIfNotExist(html_path))

	// Server
	addr := ":" + strconv.Itoa(port)
	log.Printf("Serving %s at %s", html_path, addr)
	h := http.FileServer(http.Dir(html_path))
	hs := middleware.NewHttpLogger(
		middleware.NewUpdateData(
			h, updateJsonReports(json_path, html_path)))
	http.ListenAndServe(addr, hs)
}
