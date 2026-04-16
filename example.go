// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// func main() {
// 	text := "Flight from #LAX to #JFK"

// 	airports := map[string]string{
// 		"LAX": "Los Angeles International Airport",
// 		"JFK": "John F Kennedy International Airport",
// 	}

// 	re := regexp.MustCompile(`#[A-Z]{3}`)
// 	matches := re.FindAllString(text, -1)

// 	for _, match := range matches {
// 		code := match[1:]
// 		name, found := airports[code]
// 		text = strings.ReplaceAll(text,match, name)
// 		fmt.Println("match:", match)
// 		fmt.Println("code:", code)
// 		fmt.Println("found:", found)
// 		fmt.Println("name:", name)
// 		fmt.Println(text)
// 	}
// }