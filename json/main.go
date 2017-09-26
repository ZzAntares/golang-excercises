package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ttacon/chalk"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Bar struct {
	Date                   time.Time
	Open, High, Low, Close float64
}

type Columns map[string]int

func (col Columns) fit(record []string) Bar {
	open, _ := strconv.ParseFloat(record[col["Open"]], 64)
	high, _ := strconv.ParseFloat(record[col["High"]], 64)
	low, _ := strconv.ParseFloat(record[col["Low"]], 64)
	close, _ := strconv.ParseFloat(record[col["Close"]], 64)

	date, _ := time.Parse("2006-01-02", record[col["Date"]])

	return Bar{
		Date:  date,
		Open:  open,
		High:  high,
		Low:   low,
		Close: close,
	}
}

func parseHeaders(headers []string) Columns {
	var columns Columns = make(map[string]int)

	for i, header := range headers {
		columns[header] = i
	}

	return columns
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: csv2json <input.csv> <output.json>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	headers, err := csvReader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	var columns Columns = parseHeaders(headers)
	var bars []Bar = make([]Bar, 0)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		var bar Bar = columns.fit(record)
		bars = append(bars, bar)
	}

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer outputFile.Close()

	json.NewEncoder(outputFile).Encode(bars)
	fmt.Println(chalk.Green, "Done!", chalk.ResetColor)
}
