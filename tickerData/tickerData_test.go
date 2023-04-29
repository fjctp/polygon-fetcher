package tickerData_test

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/fjctp/polygon-fetcher/tickerData"
	"github.com/fjctp/polygon-fetcher/utils"
	"github.com/polygon-io/client-go/rest/models"
)

func getTestData() (string, string, []models.StockFinancial,
	[]models.Dividend) {

	// Reference data
	finRef := models.StockFinancial{
		CompanyName: "APPLE", StartDate: "01-01-1990"}
	dividendRef := models.Dividend{
		Ticker: "AAPL", CashAmount: 1234}

	// Output data
	const ticker = "AAPL"
	const company_name = "Apple"
	finData := make([]models.StockFinancial, 2)
	finData[0] = finRef
	finData[1] = finRef

	dividendData := make([]models.Dividend, 2)
	dividendData[0] = dividendRef
	dividendData[1] = dividendRef

	return ticker, company_name, finData, dividendData
}

// TestNew calls tickerData.New, checking for a valid return value.
func TestNew(t *testing.T) {
	ticker, cname, fdata, ddata := getTestData()
	testData := tickerData.New(ticker, cname, fdata, ddata)

	// Check TickerData.Ticker
	val1 := testData.Ticker
	want1 := ticker
	if val1 != want1 {
		t.Fatalf(`Ticker = %s, want match for %s`, val1, want1)
	}

	// Check TickerData.Financial
	val2 := testData.Financial
	want2 := fdata
	if len(val2) != len(want2) {
		t.Fatalf(`len(testData.Financial) = %d, want match for %d`,
			len(val2), len(want2))
	}

	// Check TickerData.Dividend
	val3 := testData.Dividend
	want3 := ddata
	if len(val3) != len(want3) {
		t.Fatalf(`len(testData.Dividend) = %d, want match for %d`,
			len(val3), len(want3))
	}
}

func TestReadFile(t *testing.T) {
	ticker, cname, fdata, ddata := getTestData()
	testData := tickerData.New(ticker, cname, fdata, ddata)

	// Make temporary directory
	out_dir, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(out_dir)

	// Write data
	err = testData.Write(out_dir)
	if err != nil {
		t.Fatalf("Cannot write to a JSON file!")
	}

	file_path := path.Join(out_dir, ticker+".json")
	read_data, err := tickerData.ReadFile(file_path)
	if err != nil {
		t.Fatalf("Cannot read JSON file!")
	}

	// Check TickerData.Ticker
	val1 := read_data.Ticker
	want1 := ticker
	if val1 != want1 {
		t.Fatalf(`Ticker = %s, want match for %s`, val1, want1)
	}

	// Check TickerData.Financial
	val2 := read_data.Financial
	want2 := fdata
	if len(val2) != len(want2) {
		t.Fatalf(`len(read_data.Financial) = %d, want match for %d`,
			len(val2), len(want2))
	}

	// Check TickerData.Dividend
	val3 := read_data.Dividend
	want3 := ddata
	if len(val3) != len(want3) {
		t.Fatalf(`len(read_data.Dividend) = %d, want match for %d`,
			len(val3), len(want3))
	}
}

// TestWrite calls TickerData.Write with a temporary directory,
// checking for the generated JSON data file.
func TestWrite(t *testing.T) {
	ticker, cname, fdata, ddata := getTestData()
	testData := tickerData.New(ticker, cname, fdata, ddata)

	// Make temporary directory
	out_dir, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(out_dir)

	// Write data
	err = testData.Write(out_dir)
	if err != nil {
		t.Fatalf("Cannot write to a JSON file!")
	}

	// Check the data file
	file_path := path.Join(out_dir, ticker+".json")
	if !utils.Exist(file_path) {
		t.Fatalf(`%s does not exist`, file_path)
	}
}
