package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	result = 0
)

func remove(grid *[][]rune) bool {
	backup := fmt.Sprint(grid)

	maxX := len((*grid)[0]) - 1
	maxY := len(*grid) - 1

	for y, row := range *grid {
		for x := range row {
			if (*grid)[y][x] != '@' {
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

					if (*grid)[y+yOffset][x+xOffset] == '@' {
						adjacent++
					}
				}
			}

			if adjacent < 4 {
				(*grid)[y][x] = ' '
				result++
			}
		}
	}

	return backup != fmt.Sprint(grid)
}

func solve(input string) int {
	rows := strings.Split(input, "\n")
	var grid [][]rune
	for y, row := range rows {
		grid = append(grid, make([]rune, len(row)))
		for x, r := range []rune(row) {
			grid[y][x] = r
		}
	}

	for remove(&grid) {
	}

	return result
}

func main() {
	input := os.Args[1]
	fmt.Println(solve(input))
}
