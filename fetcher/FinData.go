package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/polygon-io/client-go/rest/models"
)

type FinData struct {
	Ticker string                  `json:"ticker"`
	Data   []models.StockFinancial `json:"data"`
}

// FIXME: not working
// error message: json: cannot unmarshal object into Go value of type []models.StockFinancial
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
