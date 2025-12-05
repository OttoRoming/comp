package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(input string) int {
	rows := strings.Split(input, "\n")
	maxX := len(rows[0]) - 1
	maxY := len(rows) - 1

	result := 0

	for y, row := range rows {
		for x := range row {
			if rows[y][x] != '@' {
				continue
			}

			adjacent := 0
			for yOffset := -1; yOffset <= 1 && adjacent < 4; yOffset++ {
				for xOffset := -1; xOffset <= 1 && adjacent < 4; xOffset++ {
					if yOffset == 0 && xOffset == 0 {
						continue
					}
					if x+xOffset < 0 || x+xOffset > maxX {
						continue
					}
					if y+yOffset < 0 || y+yOffset > maxY {
						continue
					}

					if rows[y+yOffset][x+xOffset] == '@' {
						adjacent++
					}
				}
			}

			if adjacent < 4 {
				result++
			}
		}
	}

	return result
}

func main() {
	input := os.Args[1]
	fmt.Println(solve(input))
}
