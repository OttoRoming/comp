package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func largestJoltage(input string) int {
	var digits []int
	for _, c := range input {
		digit, _ := strconv.Atoi(string(c))
		digits = append(digits, digit)
	}

	max := -1
	maxI := -1
	for i, digit := range digits[:len(digits)-1] {
		if digit > max {
			max = digit
			maxI = i
		}
	}

	maxSecond := -1
	for _, digit := range digits[maxI+1:] {
		if digit > maxSecond {
			maxSecond = digit
		}
	}

	return max*10 + maxSecond
}

func solve(input string) int {
	result := 0

	for r := range strings.SplitSeq(input, "\n") {
		result += largestJoltage(r)
	}

	return result
}

func main() {
	input := os.Args[1]
	fmt.Println(solve(input))
}
