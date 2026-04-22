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

	expected := []string{"name", "iso_country", "municipality", "icao_code", "iata_code", "coordinates"}
	for i, col := range expected {
		if records[0][i] != col {
			return nil, nil, errors.New("bad header")
		}
	}

	iataMap := make(map[string]Airport)
	icaoMap := make(map[string]Airport)

	for _, row := range records[1:] {
		if len(row) < 6 {
			return nil, nil, errors.New("bad row")
		}

		for i := 0; i < 6; i++ {
			if row[i] == "" {
				return nil, nil, errors.New("blank cell")
			}
		}

		name := row[0]
		city := row[2]
		icao := row[3]
		iata := row[4]

		a := Airport{
			Name: name,
			City: city,
		}
		iataMap[iata] = a
		icaoMap[icao] = a
	}
	return iataMap, icaoMap, nil
}
