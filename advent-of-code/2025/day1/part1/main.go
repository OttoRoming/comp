package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var dial int64 = 50
	result := 0

	// lines := strings.SplitSeq(os.Args[1], "\n")
	for line := range strings.SplitSeq(os.Args[1], "\n") {
		direction := line[0]
		value, _ := strconv.ParseInt(line[1:], 10, 64)

		if direction == 'L' {
			value *= -1
		}
		dial += value

		// wrap around
		if dial < 100 {
			dial += 100
		}
		dial %= 100

		if dial == 0 {
			result++
		}
	}

	fmt.Println(result)
}
