package main

import "strings"

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	cyan   = "\033[36m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	purple = "\033[35m"
	gray   = "\033[90m"
	bold   = "\033[1m"
	italic = "\033[3m"
)

func color(s string, enable bool, c ...string) string {
	if !enable {
		return s
	}
	return strings.Join(c, "") + s + reset
}
