package main

import (
	"fmt"
	"os"
)

const (
	usageText = `itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv

itinerary usage with extra features:
go run . ./input.txt ./output.txt ./airport-lookup.csv on`
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-h" {
		fmt.Println(usageText)
		return
	}

	bonusMode := false

	if len(os.Args) == 5 {
		if os.Args[4] != "on" && os.Args[4] != "off" {
			fmt.Println(usageText)
			return
		}
		bonusMode = os.Args[4] == "on"
	} else if len(os.Args) != 4 {
		fmt.Println(usageText)
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]
	airportPath := os.Args[3]

	if !fileExists(inputPath) {
		fmt.Println("Input not found")
		return
	}

	if !fileExists(airportPath) {
		fmt.Println("Airport lookup not found")
		return
	}

	iataMap, icaoMap, err := loadAirports(airportPath)
	if err != nil {
		fmt.Println("Airport lookup malformed")
		return
	}

	text, err := readTextFile(inputPath)
	if err != nil {
		fmt.Println("Input not found")
		return
	}

	normalized := normalizeVerticalWhitespace(text)
	normalized = trimExtraBlankLines(normalized)

	output := replaceAirportCodes(normalized, iataMap, icaoMap, false)
	output = replaceDateTimes(output, false)

	err = writeTextFile(outputPath, output)
	if err != nil {
		return
	}

	if bonusMode {
		preview := replaceAirportCodes(normalized, iataMap, icaoMap, true)
		preview = replaceDateTimes(preview, true)

		fmt.Println(preview)
	}
}
