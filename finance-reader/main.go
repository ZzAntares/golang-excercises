package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Bar struct {
	Date   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
}

func parseHeaders(record []string) map[string]int {
	var columns map[string]int = make(map[string]int)

	for idx, column := range record {
		columns[column] = idx
	}

	return columns
}

func parseBarRecord(columns map[string]int, record []string) Bar {
	open, _ := strconv.ParseFloat(record[columns["Open"]], 64)
	high, _ := strconv.ParseFloat(record[columns["High"]], 64)
	low, _ := strconv.ParseFloat(record[columns["Low"]], 64)
	close, _ := strconv.ParseFloat(record[columns["Close"]], 64)
	volume, _ := strconv.Atoi(record[columns["Volume"]])

	return Bar{
		Date:   record[columns["Date"]],
		Open:   open,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: volume,
	}
}

func main() {
	var filename string = "spy.csv"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	headers, err := csvReader.Read()
	if err != nil {
		log.Fatalln(headers)
	}

	var columns map[string]int = parseHeaders(headers)

	var bars []Bar = make([]Bar, 0)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		bar := parseBarRecord(columns, record)
		bars = append(bars, bar)
	}

	fmt.Println(bars)
}
