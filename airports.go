package main

import (
	"encoding/csv"
	"errors"
	"os"
)

type Airport struct {
	Name string
	City string
}

func loadAirports(airportPath string) (map[string]Airport, map[string]Airport, error) {
	file, err := os.Open(airportPath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	if len(records) < 1 {
		return nil, nil, errors.New("empty csv")
	}

	header := records[0]
	headerMap := make(map[string]int)

	for i, col := range header {
		headerMap[col] = i
	}

	expected := []string{"name", "iso_country", "municipality", "icao_code", "iata_code", "coordinates"}
	for _, col := range expected {
		if _, ok := headerMap[col]; !ok {
			return nil, nil, errors.New("missing expected column: " + col)
		}
	}

	iataMap := make(map[string]Airport)
	icaoMap := make(map[string]Airport)

	for _, row := range records[1:] {
		if len(row) < len(header) {
			return nil, nil, errors.New("bad row")
		}

		for i := 0; i < 6; i++ {
			if row[i] == "" {
				return nil, nil, errors.New("blank cell")
			}
		}

		name := row[headerMap["name"]]
		city := row[headerMap["municipality"]]
		icao := row[headerMap["icao_code"]]
		iata := row[headerMap["iata_code"]]

		a := Airport{
			Name: name,
			City: city,
		}
		iataMap[iata] = a
		icaoMap[icao] = a
	}
	return iataMap, icaoMap, nil
}
