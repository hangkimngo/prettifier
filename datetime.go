package main

import (
	"regexp"
	"strings"
	"time"
)

func formatDateTime(kind string, value string) (string, bool) {
	var t time.Time
	var err error

	if strings.HasSuffix(value, "Z") {
		t, err = time.Parse("2006-01-02T15:04Z", value)
	} else {
		t, err = time.Parse("2006-01-02T15:04-07:00", value)
	}

	if err != nil {
		return "", false
	}

	offset, ok := formatOffset(value)
	if !ok {
		return "", false
	}

	switch kind {
	case "D":
		return t.Format("02 Jan 2006"), true
	case "T12":
		return t.Format("03:04PM") + " " + offset, true
	case "T24":
		return t.Format("15:04") + " " + offset, true
	default:
		return "", false
	}

}

func formatOffset(value string) (string, bool) {
	if strings.HasSuffix(value, "Z") {
		return "(+00:00)", true
	}

	if len(value) < 6 {
		return "", false
	}

	tz := value[len(value)-6:]

	matched, _ := regexp.MatchString(`^[+-]\d{2}:\d{2}$`, tz)
	if !matched {
		return "", false
	}
	return "(" + tz + ")", true
}

func replaceDateTimes(text string) string {
	re := regexp.MustCompile(`(D|T12|T24)\(([^)]+)\)`)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}

		kind := parts[1]
		value := parts[2]

		formatted, ok := formatDateTime(kind, value)
		if !ok {
			return match
		}

		return formatted
	})
}
