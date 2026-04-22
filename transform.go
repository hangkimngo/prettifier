package main

import (
	"regexp"
	"strings"
)

func replaceAirportCodes(text string, iataMap, icaoMap map[string]Airport) string {
	re := regexp.MustCompile(`\*?##[A-Z]{4}|\*?#[A-Z]{3}`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		wantCity := false
		original := match

		if strings.HasPrefix(match, "*") {
			wantCity = true
			match = match[1:]
		}

		var a Airport
		var found bool

		if strings.HasPrefix(match, "##") {
			code := match[2:]
			a, found = icaoMap[code]
		} else {
			code := match[1:]
			a, found = iataMap[code]
		}

		if !found {
			return original
		}

		if wantCity {
			if a.City != "" {
				return a.City
			}
			return original
		}

		if a.Name != "" {
			return a.Name
		}
		return original
	})
	return text
}

func normalizeVerticalWhitespace(text string) string {
	re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000B}\x{000C}]|(\\r|\\v|\\f)`)

	text = re.ReplaceAllString(text, "\n")

	// text = strings.ReplaceAll(text, "\r\n", "\n")
	// text = strings.ReplaceAll(text, "\r", "\n")
	// text = strings.ReplaceAll(text, "\v", "\n")
	// text = strings.ReplaceAll(text, "\f", "\n")

	// text = strings.ReplaceAll(text, "\\r", "\n")
	// text = strings.ReplaceAll(text, "\\v", "\n")
	// text = strings.ReplaceAll(text, "\\f", "\n")

	return text
}

func trimExtraBlankLines(text string) string {
	extraBlanklines := regexp.MustCompile(`\n{3,}`)
	text = extraBlanklines.ReplaceAllString(text, "\n\n")
	return text
}

//
