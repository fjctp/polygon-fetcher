package report

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/fjctp/polygon-fetcher/tickerData"
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
func New(tData tickerData.TickerData, output_dir string) error {
	// get a new template
	t, err := template.New("report").Parse(templateStr)
	if err != nil {
		return err
	}

	// create data structure for the template
	r := Report{
		Name: tData.CompanyName + " Financial Report",
		Canvases: []Canvas{
			NewBalanceSheetCanvas(tData.Financial, "chart1"),
			NewCashFlowCanvas(tData.Financial, "chart2"),
			NewEpsCanvas(tData.Financial, "chart3"),
			NewIncomeProfitCanvas(tData.Financial, "chart4"),
			NewDividendCanvas(tData.Dividend, "chart5"),
		},
	}

	// output the html report with data embedded
	p := filepath.Join(output_dir, tData.Ticker+".html")
	log.Printf("Write report to %s\n", p)
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, r)
}
