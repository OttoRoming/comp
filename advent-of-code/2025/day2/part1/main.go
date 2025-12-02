package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkNumber(i int) int {
	if math.Mod(math.Floor(math.Log10(float64(i))), 2) == 0.0 {
		return 0
	}

	str := fmt.Sprint(i)
	first := str[:len(str)/2]
	if strings.Repeat(first, 2) == str {
		return i
	}
	return 0
}

func checkRange(result *int, r string) {
	split := strings.Split(r, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	for i := start; i <= end; i++ {
		*result += checkNumber(i)
	}
}

func solve(input string) int {
	var result int = 0

	for r := range strings.SplitSeq(input, ",") {
		checkRange(&result, r)
	}

	return result
}

func main() {
	input := os.Args[1]
	fmt.Println(solve(input))
}
