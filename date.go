// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	text := "Depart: D(2007-04-05T12:30-02:00)\nBoard: T12(2007-04-05T12:30-02:00)"

// 	re := regexp.MustCompile(`(D|T12|T24)\(([^)]+)\)`)
// 	matches := re.FindAllStringSubmatch(text, -1)

// 	fmt.Println(matches)

//		for _, m := range matches {
//			fmt.Println("full:", m[0])
//			fmt.Println("kind:", m[1])
//			fmt.Println("value:", m[2])
//			fmt.Println()
//		}
//	}
package main

import (
	"fmt"
	"time"
)

func main() {
	value := "2007-04-05T12:30-02:00"

	t, err := time.Parse("2006-01-02T15:04-07:00", value)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	fmt.Println(t)
	fmt.Println(t.Format("02 Jan 2006"))
	fmt.Println(t.Format("03:04PM"))
	fmt.Println(t.Format("15:04"))
}
