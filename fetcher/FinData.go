package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/polygon-io/client-go/rest/models"
)

type FinData struct {
	Ticker string
	Data   []models.StockFinancial
}

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
