package main

import (
	"regexp"
	"strings"
	"time"
)

func formatDateTime(kind string, value string, bonusMode bool) (string, bool) {
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

	offset, ok := formatOffset(value, bonusMode)
	if !ok {
		return "", false
	}

	switch kind {
	case "D":
		return color(t.Format("02 Jan 2006"), blue, bonusMode), true
	case "T12":
		return color(t.Format("03:04PM")+" "+offset, yellow, bonusMode), true
	case "T24":
		return color(t.Format("15:04")+" "+offset, yellow, bonusMode), true
	default:
		return "", false
	}

}

func formatOffset(value string, colorMode bool) (string, bool) {
	if strings.HasSuffix(value, "Z") {
		return color("(+00:00)", purple, colorMode), true
	}

	if len(value) < 6 {
		return "", false
	}

	tz := value[len(value)-6:]

	matched, _ := regexp.MatchString(`^[+-]\d{2}:\d{2}$`, tz)
	if !matched {
		return "", false
	}
	return color("("+tz+")", purple, colorMode), true
}

func replaceDateTimes(text string, colorMode bool) string {
	re := regexp.MustCompile(`(D|T12|T24)\(([^)]+)\)`)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}

		kind := parts[1]
		value := parts[2]

		formatted, ok := formatDateTime(kind, value, colorMode)
		if !ok {
			return match
		}

		return formatted
	})
}
