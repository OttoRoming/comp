package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int64) int64 {
	if x < 0 {
		x *= -1
	}
	return x
}

func main() {
	var dial int64 = 50
	var result int64 = 0

	// lines := strings.SplitSeq(os.Args[1], "\n")
	for line := range strings.SplitSeq(os.Args[1], "\n") {
		dir_char := line[0]
		var dir int64 = 1
		value, _ := strconv.ParseInt(line[1:], 10, 64)

		if dir_char == 'L' {
			dir = -1
		}

		for range value {
			dial += dir
			dial += 1000000
			dial %= 100

			if dial == 0 {
				result++
			}
		}
	}

	fmt.Println(result)
}
