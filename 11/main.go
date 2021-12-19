package main

import (
	"fmt"
	"math"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

var length, width int

func outOfBounds(i, j int) bool {
	return (i < 0 || i >= length || j < 0 || j >= width)
}

func gridHasGT9(grid [][]int) bool {
	for _, line := range grid {
		for _, number := range line {
			if number > 9 {
				return true
			}
		}
	}
	return false
}

func main() {
	grid := utils.ReadFileGridInts("11/input.txt")

	length = len(grid)
	width = len(grid[0])

	var flashes int

	for i := 0; true; i++ {
		var newGrid [][]int
		for _, line := range grid {
			var newLine []int
			for _, number := range line {
				newLine = append(newLine, number+1)
			}
			newGrid = append(newGrid, newLine)
		}

		for gridHasGT9(newGrid) {
			grid = newGrid

			for i, line := range grid {
				for j, number := range line {
					if number > 9 {
						for _, diff := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
							i_diff, j_diff := i+diff[0], j+diff[1]
							if !outOfBounds(i_diff, j_diff) {
								// Just incrementing works here, because newGrid==grid
								newGrid[i_diff][j_diff]++
								// Mark this as an octopus that has flashed
							}
						}
						newGrid[i][j] = math.MinInt32
						flashes++
					}
				}
			}
		}

		grid = newGrid

		var flashedThisTime int
		// Set all the flashed octopuses to 0
		for i, line := range grid {
			for j, num := range line {
				if num < 0 {
					newGrid[i][j] = 0
					flashedThisTime++
				}
			}
		}

		// Part 1
		if i == 99 {
			fmt.Println(flashes)
		}

		// Part 2
		if flashedThisTime == width*length {
			fmt.Println(i + 1)
			return
		}

	}

}
