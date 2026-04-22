package main

import (
	"regexp"
	"strings"
)

const (
	reset  = "\033[0m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	purple = "\033[35m"
	gray   = "\033[90m"
)

func formatOutput(text string) string {
	reISO := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}(?::\d{2})?(Z|[+-]\d{2}:\d{2})`)

	text = reISO.ReplaceAllStringFunc(text, func(match string) string {
		// split parts
		date := match[:10] // YYYY-MM-DD
		rest := match[10:] // T...

		// split T
		timeAndOffset := rest[1:] // bỏ T
		T := "T"

		// tìm offset
		var timePart, offsetPart string

		if strings.HasSuffix(timeAndOffset, "Z") {
			timePart = timeAndOffset[:len(timeAndOffset)-1]
			offsetPart = "Z"
		} else {
			// tìm + hoặc -
			i := strings.HasPrefix(timeAndOffset, "+-")
			if i {
				timePart = timeAndOffset[:i]
				offsetPart = timeAndOffset[i:]
			} else {
				timePart = timeAndOffset
			}
		}

		// build colored string
		result := blue + date + reset +
			gray + T + reset +
			yellow + timePart + reset

		if offsetPart != "" {
			result += purple + offsetPart + reset
		}

		return result
	})

	return text
}
