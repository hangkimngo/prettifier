package main

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	cyan   = "\033[36m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	purple = "\033[35m"
	gray   = "\033[90m"
)

func color(s, c string, enable bool) string {
	if !enable {
		return s
	}
	return c + s + reset
}
