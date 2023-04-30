package report

// A canvas that contains a chart.
type Canvas struct {
	Id   string
	Data Chart
}

// A chart that contains data and chart options. See chart.js for details
type Chart struct {
	Type    string       `json:"type"`
	Data    ChartData    `json:"data"`
	Options ChartOptions `json:"options"`
}

type ChartDataSet struct {
	Label       string    `json:"label"`
	Data        []float32 `json:"data"`
	BorderWidth int       `json:"borderWidth"`
}

type ChartData struct {
	Labels   []string       `json:"labels"`
	Datasets []ChartDataSet `json:"datasets"`
}

type ChartOptions struct {
}
