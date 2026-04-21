package main

import (
	"fmt"
	"os"
)

const usageText = `itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv`

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-h" {
		fmt.Println(usageText)
		return
	}

	if len(os.Args) != 4 {
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

	text = normalizeVerticalWhitespace(text)
	text = trimExtraBlankLines(text)
	text = replaceAirportCodes(text, iataMap, icaoMap)
	text = replaceDateTimes(text)

	err = writeTextFile(outputPath, text)
	if err != nil {
		return
	}

}
