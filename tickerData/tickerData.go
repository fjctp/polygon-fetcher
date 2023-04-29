package tickerData

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/polygon-io/client-go/rest/models"
)

type TickerData struct {
	Ticker    string                  `json:"ticker"`
	Financial []models.StockFinancial `json:"financial"`
	Dividend  []models.Dividend       `json:"dividend"`
}

func New(ticker string, fin []models.StockFinancial,
	dividend []models.Dividend) TickerData {
	return TickerData{ticker, fin, dividend}
}

// Read data from a JSON file and output a populated TickerData structure
func ReadFile(path string) (TickerData, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return TickerData{}, err
	}

	var d TickerData
	err = json.Unmarshal(content, &d)
	if err != nil {
		return TickerData{}, err
	}
	return d, nil
}

// Write TickerData to a JSON file
func (d TickerData) Write(out_dir string) error {
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
