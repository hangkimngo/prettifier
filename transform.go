package main

import (
	"regexp"

)


func replaceAirportCodes(text string, iataMap, icaoMap map[string]string) string {
	reICAO := regexp.MustCompile(`##[A-Z]{4}`)
	text= reICAO.ReplaceAllStringFunc(text, func(match string) string {
		code := match[2:]
		name, found := icaoMap[code]
		if found {
			return name
		}
		return match
	})

	reIATA := regexp.MustCompile(`(^|[^#])(#[A-Z]{3})`)
	text = reIATA.ReplaceAllStringFunc(text, func(match string) string {
		prefix:=""
		token:=match

		if len(match) > 4 {
			prefix = match[:1]
			token = match[1:]
		}
		
		code:=token[1:]
		name, found := iataMap[code]
		if found {
			return prefix + name
		}
		return match
	})

	return text
	
}