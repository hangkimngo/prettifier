package main

import (
	"regexp"
	"strings"
)

func replaceAirportCodes(text string, iataMap, icaoMap map[string]Airport, bonusMode bool) string {
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

		if wantCity && bonusMode {
			if a.City != "" {
				return color(a.City, green, bonusMode)
			}
			return original
		}

		if a.Name != "" {
			return color(a.Name, green, bonusMode)
		}
		return original
	})
	return text
}

func normalizeVerticalWhitespace(text string) string {
	// re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000B}\x{000C}]|(\\r|\\v|\\f)`)
	re := regexp.MustCompile(`\r\n|\r|\n|\v|\f|\\r|\\v|\\f`)
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
	extraBlanklines := regexp.MustCompile(`\n[ \t]*\n(?:[ \t]*\n)+`)
	text = extraBlanklines.ReplaceAllString(text, "\n\n")
	return text
}
