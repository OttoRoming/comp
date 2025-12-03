package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	batteryCount = 12
)

func max(digits []int) (int, int) {
	max := -1
	index := -1
	for i, digit := range digits {
		if digit > max {
			max = digit
			index = i
		}
	}

	return max, index
}

func largestJoltage(digits []int, batteries int) int {
	result := 0

	if batteries <= 0 {
		return 0
	}

	r := len(digits) - batteries
	slice := digits[:r+1]
	max, i := max(slice)
	if batteries > 0 {
		result += largestJoltage(digits[i+1:], batteries-1)
	}

	result += max * int(math.Pow10(batteries))
	return result
}

func solve(input string) int {
	result := 0

	for r := range strings.SplitSeq(input, "\n") {
		var digits []int
		for _, c := range r {
			digit, _ := strconv.Atoi(string(c))
			digits = append(digits, digit)
		}

		bugged := largestJoltage(digits, batteryCount)
		result += bugged / 10
	}

	return result
}

func main() {
	input := os.Args[1]
	fmt.Println(solve(input))
}
