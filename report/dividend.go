package report

import (
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

// Create a canvas for data for dividend
func NewDividendCanvas(data []models.Dividend, id string) Canvas {
	chart := newDividendChart(data)
	return Canvas{Id: id, Data: chart}
}

// Create a chart for the balance sheet data
func newDividendChart(data []models.Dividend) Chart {
	// extract date and information from data
	dsLabel := "Dividend" // labels for lines
	var xdata []string    // xdata for a chart
	var ydata []float32   // ydata for a chart

	for _, record := range data {
		// get fiscal year for each data point
		// insert the new data at the start of the slice
		dDate := time.Time(record.DeclarationDate)
		fiscalLabel := dDate.Format("2006-01-02")
		xdata = append(xdata, fiscalLabel)
		ydata = append(ydata, float32(record.CashAmount))

	}

	// create an array of dataset object
	var datasets []ChartDataSet
	cds := ChartDataSet{
		Label:       dsLabel,
		Data:        ydata,
		BorderWidth: 1,
	}
	datasets = append(datasets, cds)

	// Create data struct
	cData := ChartData{
		Labels:   xdata,
		Datasets: datasets,
	}
	cOpt := ChartOptions{}
	return Chart{Type: "line", Data: cData, Options: cOpt}
}
