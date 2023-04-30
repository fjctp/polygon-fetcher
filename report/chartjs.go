package report

// Create a new Canvas
func NewCanvas(id string, c Chart) Canvas {
	return Canvas{id, c}
}

// Create a new Chart with type = line
func NewLineChart(xLabel []string, yLabel map[string][]float32,
	dsLabel map[string]string) Chart {

	datasets := make([]ChartDataSet, 0, 4)
	for key, val := range yLabel {
		lineLabel := dsLabel[key]
		ds := newChartDataSet(lineLabel, val, 1)
		datasets = append(datasets, ds)
	}
	cData := newChartData(xLabel, datasets)

	cOpt := ChartOptions{}
	return Chart{Type: "line", Data: cData, Options: cOpt}
}

// Create a new ChartDataSet
func newChartDataSet(label string, data []float32,
	borderWidth int) ChartDataSet {
	return ChartDataSet{label, data, borderWidth}
}

// Create a new ChartData
func newChartData(xLabels []string, dataset []ChartDataSet) ChartData {
	return ChartData{xLabels, dataset}
}

// A canvas that contains a chart. See chart.js for details
type Canvas struct {
	Id   string
	Data Chart
}

// A chart that contains data and chart options.
type Chart struct {
	Type    string       `json:"type"`
	Data    ChartData    `json:"data"`
	Options ChartOptions `json:"options"`
}

// A chart dataset that contains the label of the set, and y-axis data
type ChartDataSet struct {
	Label       string    `json:"label"` // label for the dataset
	Data        []float32 `json:"data"`  // y-axis data
	BorderWidth int       `json:"borderWidth"`
}

// A chart data contains x-axis label and multiple y-axis datasets
type ChartData struct {
	Labels   []string       `json:"labels"`   // x-axis label
	Datasets []ChartDataSet `json:"datasets"` // y-axis datasets
}

// A chart options for Chart type
type ChartOptions struct {
}
