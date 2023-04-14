package report

import "github.com/fjctp/polygon-fetcher/fetcher"

// A FinanicalChartPair that constains data for a finanical chart
type FinanicalChartPair struct {
	Id     string   // Canvas ID
	Name   string   // finanical statement name
	Fields []string // field names from the financial statement
}

// Get FinanicalChartPair
func NewFinanicalChartPair(id string, name string, fields ...string) FinanicalChartPair {
	return FinanicalChartPair{Id: id, Name: name, Fields: fields}
}

// Create a canvas for the balance sheet data
func getFinancialCanvas(finData fetcher.FinData, pair FinanicalChartPair) Canvas {
	chart := getFinancialChart(finData, pair)
	return Canvas{Id: pair.Id, Data: chart}
}

// Create a chart for the balance sheet data
func getFinancialChart(finData fetcher.FinData, pair FinanicalChartPair) Chart {
	// extract data from finData
	datasetLabels := make(map[string]string) // labels for lines
	var xdata []string                       // xdata for a chart
	ydata := make(map[string][]float32)
	for i, record := range finData.Data {
		// get fiscal year for each data point
		// insert the new data at the start of the slice
		fiscalLabel := record.FiscalYear + " " + record.FiscalPeriod
		xdata = append([]string{fiscalLabel}, xdata...)

		// get data for each keys
		statement := record.Financials[pair.Name]
		for _, key := range pair.Fields {
			info := statement[key]

			// get label for the line
			if i == 0 {
				datasetLabels[key] = info.Label
			}
			// get ydata for each data point
			// insert the new data at the start of the slice
			val := float32(info.Value)
			ydata[key] = append([]float32{val}, ydata[key]...)
		}
	}

	// create an array of dataset object
	var datasets []ChartDataSet
	for _, key := range pair.Fields {
		cds := ChartDataSet{
			Label:       datasetLabels[key],
			Data:        ydata[key],
			BorderWidth: 1,
		}
		datasets = append(datasets, cds)
	}

	// Create data struct
	cData := ChartData{
		Labels:   xdata,
		Datasets: datasets,
	}
	cOpt := ChartOptions{}
	return Chart{Type: "line", Data: cData, Options: cOpt}
}
