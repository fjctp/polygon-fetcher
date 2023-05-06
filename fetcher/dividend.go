package fetcher

import (
	"context"
	"log"
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

// Fetch dividend data from Polygon using its client and
// save the data in JSON format
func FetchDividendData(f Fetcher, ticker string,
	count int) ([]models.Dividend, error) {

	params := models.ListDividendsParams{}.
		WithTicker(models.EQ, ticker).
		WithDividendType(models.DividendCD).
		WithLimit(count).
		WithOrder(models.Asc).
		WithSort("declaration_date")
	iter := f.polygon.ListDividends(
		context.Background(), params)

	// read the next record and keep the data in a slice
	var data []models.Dividend
	for ind := 0; ind < count; ind++ {
		if !iter.Next() {
			if ind == 0 {
				return data, nil
			}
			// reach the end of the results
			log.Print("End of record")
			break
		}

		item := iter.Item()
		declarationDate := time.Time(item.DeclarationDate)
		log.Printf("%d, %s\n", ind, declarationDate.Format("2006-01-02"))

		data = append(data, item)
	}

	return data, iter.Err()
}
