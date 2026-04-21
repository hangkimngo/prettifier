package main

import (
	"regexp"
)

func replaceAirportCodes(text string, iataMap, icaoMap map[string]string) string {
	reICAO := regexp.MustCompile(`##[A-Z]{4}`)
	text = reICAO.ReplaceAllStringFunc(text, func(match string) string {
		code := match[2:]
		name, found := icaoMap[code]
		if found {
			return name
		}
		return match
	})

	reIATA := regexp.MustCompile(`(^|[^#])(#[A-Z]{3})`)
	text = reIATA.ReplaceAllStringFunc(text, func(match string) string {
		prefix := ""
		token := match

		if len(match) > 4 {
			prefix = match[:1]
			token = match[1:]
		}

		code := token[1:]
		name, found := iataMap[code]
		if found {
			return prefix + name
		}
		return match
	})
	return text
}

func normalizeVerticalWhitespace(text string) string {
	re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000B}\x{000C}]|(\\r|\\v|\\f)`)

	text = re.ReplaceAllString(text, "\n")
	return text
}

func trimExtraBlankLines(text string) string {
	extraBlanklines := regexp.MustCompile(`\n{3,}`)
	text = extraBlanklines.ReplaceAllString(text, "\n\n")
	return text
}

//
