package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type State struct {
	id               int
	name             string
	abbreviation     string
	country          string
	censusRegionName string
}

func (s *State) ToString() string {
	return fmt.Sprintf(
		"ID: %d\nName: %s\nAbbreviation: %s\nCountry: %s\nRegion: %s\n",
		s.id, s.name, s.abbreviation, s.country, s.censusRegionName,
	)
}

type StateDict map[string]State

func parseStateRecord(columns map[string]int, record []string) (*State, error) {
	stateId, err := strconv.Atoi(record[columns["id"]])

	if err != nil {
		return nil, err
	}

	var state State = State{
		id:               stateId,
		name:             record[columns["name"]],
		abbreviation:     record[columns["abbreviation"]],
		country:          record[columns["country"]],
		censusRegionName: record[columns["census_region_name"]],
	}

	return &state, nil
}

func parseHeadersMap(headers []string) map[string]int {
	var columns map[string]int = make(map[string]int)

	for idx, column := range headers {
		columns[column] = idx
	}

	return columns
}

func loadStates(filename string) (*StateDict, error) {
	stateFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer stateFile.Close()

	// Parse csv file
	csvReader := csv.NewReader(stateFile)

	headers, err := csvReader.Read() // Discard headers

	if err != nil {
		return nil, err
	}

	columns := parseHeadersMap(headers)
	var states StateDict = make(StateDict)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return &states, err
		}

		state, err := parseStateRecord(columns, record)

		if err != nil {
			return &states, err
		}

		states[state.abbreviation] = *state
	}

	return &states, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: statable <ABBR>")
	}

	var abbr string = strings.ToUpper(os.Args[1])

	states, err := loadStates("state_table.csv")

	if err != nil {
		log.Fatalln(err)
	}

	var lookupStates StateDict = *states
	var state State = lookupStates[abbr]

	fmt.Println(state.ToString())
}
