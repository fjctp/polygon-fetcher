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
func New(finData fetcher.FinData, output_dir string) error {
	// get a new template
	t, err := template.New("report").Parse(templateStr)
	if err != nil {
		return err
	}

	// create data structure for the template
	r := Report{
		Name: finData.Data[0].CompanyName + " Financial Report",
		Canvases: []Canvas{
			NewBalanceSheetCanvas(finData, "chart1"),
			NewCashFlowCanvas(finData, "chart2"),
			NewEpsCanvas(finData, "chart3"),
			NewIncomeProfitCanvas(finData, "chart4"),
		},
	}

	// output the html report with data embedded
	p := filepath.Join(output_dir, finData.Ticker+".html")
	log.Printf("Write report to %s\n", p)
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, r)
}
