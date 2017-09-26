package main

import (
	"encoding/csv"
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var ErrNoData error = errors.New("Couldn't find requested data file.")

type Bar struct {
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume uint64    `json:"volume"`
}

func parseHeaders(headers []string) map[string]int {
	var columns map[string]int = make(map[string]int)

	for i, column := range headers {
		columns[column] = i
	}

	return columns
}

func parseBarRecord(col map[string]int, record []string) *Bar {
	date, _ := time.Parse("2006-01-02", record[col["Date"]])
	open, _ := strconv.ParseFloat(record[col["Open"]], 64)
	high, _ := strconv.ParseFloat(record[col["High"]], 64)
	low, _ := strconv.ParseFloat(record[col["Low"]], 64)
	close, _ := strconv.ParseFloat(record[col["Close"]], 64)
	volume, _ := strconv.ParseUint(record[col["Volume"]], 10, 64)

	return &Bar{
		Date:   date,
		Open:   open,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: volume,
	}
}

func LoadCsvFromFile(file io.Reader) ([]Bar, error) {
	csvReader := csv.NewReader(file)
	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	var columns map[string]int = parseHeaders(headers)
	var bars []Bar = make([]Bar, 0)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		var bar *Bar = parseBarRecord(columns, record)
		bars = append(bars, *bar)
	}

	return bars, nil
}

func LoadCsv(filename string) ([]Bar, error) {
	datafile, err := os.Open(filename)
	if err != nil {
		return nil, ErrNoData
	}
	defer datafile.Close()

	return LoadCsvFromFile(datafile)
}

func FinanceHandler(res http.ResponseWriter, req *http.Request) {
	var ticker string = strings.Split(req.RequestURI, "/")[1]

	if ticker == "" {
		io.WriteString(
			res,
			"Try to navigate to a ticker: http://localhost:9000/spy",
		)
		return
	}

	data, err := LoadCsv(ticker + ".csv")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	// data is a slice with every data row as a Bar in it
	tpl, err := template.ParseFiles("plotter.html.tmpl")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	tpl.Execute(res, data)
}

func main() {
	http.HandleFunc("/", FinanceHandler)
	http.ListenAndServe("localhost:9000", nil)
}
