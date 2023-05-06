package report

import (
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

// Create a canvas for data for dividend
func NewDividendCanvas(data []models.Dividend, id string) Canvas {
	chart := newDividendChart(data)
	return NewCanvas(id, chart)
}

// Create a chart for the balance sheet data
func newDividendChart(data []models.Dividend) Chart {
	// extract date and information from data
	key := "dividend"

	dsLabel := make(map[string]string) // labels for lines
	dsLabel[key] = "Dividend"
	var xdata []string                  // xdata for a chart
	ydata := make(map[string][]float32) // ydata for a chart

	for _, record := range data {
		// get fiscal year for each data point
		// insert the new data at the start of the slice
		dDate := time.Time(record.DeclarationDate)
		fiscalLabel := dDate.Format("2006-01-02")
		xdata = append(xdata, fiscalLabel)
		ydata[key] = append(ydata[key], float32(record.CashAmount))

	}

	return NewLineChart(xdata, ydata, dsLabel)
}
