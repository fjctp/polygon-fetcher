package report

import (
	"log"

	"github.com/polygon-io/client-go/rest/models"
)

// Create a canvas with balance sheet's charts
func NewBalanceSheetCanvas(finData []models.StockFinancial, id string) Canvas {
	return newFinancialCanvas(finData, id, "balance_sheet",
		"assets", "equity", "liabilities")
}

// Create a canvas with cash flow's charts
func NewCashFlowCanvas(finData []models.StockFinancial, id string) Canvas {
	return newFinancialCanvas(finData, id, "cash_flow_statement",
		"net_cash_flow")
}

// Create a canvas with EPS's charts
func NewEpsCanvas(finData []models.StockFinancial, id string) Canvas {
	return newFinancialCanvas(finData, id, "income_statement",
		"basic_earnings_per_share")
}

// Create a canvas with Income and profit's charts
func NewIncomeProfitCanvas(finData []models.StockFinancial, id string) Canvas {
	return newFinancialCanvas(finData, id, "income_statement",
		"cost_of_revenue", "gross_profit", "net_income_loss", "revenues")
}

// Create a canvas for data in a statements (balance_sheet,
// cash_flow_statement, income_statement)
func newFinancialCanvas(finData []models.StockFinancial, id string,
	statement string, fields ...string) Canvas {
	chart := newFinancialChart(finData, statement, fields...)
	return Canvas{Id: id, Data: chart}
}

// Create a chart for the balance sheet data
func newFinancialChart(finData []models.StockFinancial,
	statementName string, fields ...string) Chart {
	// extract date and information from finData
	datasetLabels := make(map[string]string) // labels for lines
	var xdata []string                       // xdata for a chart
	ydata := make(map[string][]float32)
	for _, record := range finData {
		// get fiscal year for each data point
		// insert the new data at the start of the slice
		fiscalLabel := record.FiscalYear + " " + record.FiscalPeriod
		xdata = append(xdata, fiscalLabel)

		// get data for each keys
		statement, ok := record.Financials[statementName]
		if !ok {
			log.Printf("%s does not exist for %s.\n",
				statementName, fiscalLabel)
			continue
		}
		for _, key := range fields {
			info, ok := statement[key]
			if !ok {
				log.Printf("%s.%s does not exist for %s.\n",
					statementName, key, fiscalLabel)
				continue
			}

			// get label for the line
			if _, ok := datasetLabels[key]; !ok {
				datasetLabels[key] = info.Label
			}
			// get ydata for each data point
			// insert the new data at the start of the slice
			val := float32(info.Value)
			ydata[key] = append(ydata[key], val)
		}
	}

	// create an array of dataset object
	var datasets []ChartDataSet
	for _, key := range fields {
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
