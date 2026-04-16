package main

import (
	"encoding/csv"
	"errors"
	"os"
)

func loadAirports(airportPath string) (map[string]string, map[string]string, error) {
	file, err := os.Open(airportPath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err:= reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	if len(records) < 1 {
	return nil, nil, errors.New("empty csv")
	}

	if len(records[0]) < 6 {
	return nil, nil, errors.New("bad header")
	}

	iataMap := make(map[string]string)
	icaoMap := make(map[string]string)

	for _, row := range records[1:] {
		if len(row) < 6 {
			return nil, nil, errors.New("bad row")
			}
		
		for i:=0 ; i < 6; i++ {
			if row[i] == "" {
				return nil, nil, errors.New("blank cell")
			}
		}
		
		name := row[0]
		icao:=row[3]
		iata:= row[4]

		iataMap[iata]=name
		icaoMap[icao]=name
	}
	return iataMap,icaoMap, nil
}

