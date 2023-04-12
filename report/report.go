package report

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/fjctp/polygon-fetcher/fetcher"
)

// HTML template
const templateStr = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Name}}</title>
	</head>
	<body>
		<div>{{.Name}}</div>
		{{range .Canvases}}
        <div>
            <canvas id="{{.Id}}"></canvas>
        </div>
		{{end}}
    </body>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <script>
		{{range .Canvases}}
        new Chart(document.getElementById("{{.Id}}"), {{.Data}});
		{{end}}
    </script>
</html>`

// Create a new html report using the finanical data provided
func New(ticker string, finData fetcher.FinData, output_dir string) error {
	// get a new template
	t, err := template.New("report").Parse(templateStr)
	if err != nil {
		return err
	}

	// create data structure for the template
	// The key is the canvas ID.
	// The first value is the financial statement name,
	//   where data comes from.
	// The second and beyond values are the field names
	//   from a financial statement.
	summary := make(map[string][]string)
	summary["chart1"] = []string{"balance_sheet",
		"assets", "equity", "liabilities"}
	summary["chart2"] = []string{"cash_flow_statement",
		"net_cash_flow"}
	summary["chart3"] = []string{"income_statement",
		"basic_earnings_per_share"}
	summary["chart4"] = []string{"income_statement",
		"cost_of_revenue", "gross_profit", "net_income_loss", "revenues"}
	count := int(1)
	var canvases []Canvas
	for k, v := range summary {
		chart := getChart(finData, v[0], v[1:])

		canvas := Canvas{Id: k, Data: chart}
		canvases = append(canvases, canvas)

		count++
	}

	r := Report{
		Name:     ticker + " Financial Report",
		Canvases: canvases,
	}

	// output the html report with data embedded
	p := filepath.Join(output_dir, ticker+".html")
	log.Printf("Write report to %s\n", p)
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, r)
}

// Create a chart for the balance sheet data
func getChart(finData fetcher.FinData, sName string, keys []string) Chart {
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
		statement := record.Financials[sName]
		for _, key := range keys {
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
	for _, key := range keys {
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
