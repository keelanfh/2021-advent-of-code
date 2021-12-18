package main

import (
	"fmt"
	"sort"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

// Set up length and width as global variables so we can do bounds checking easily
var length, width int
var basinSizes []int
var lowPoints, diffs [][2]int
var grid [][]int
var inBasinGrid [][]inBasinResult

type inBasinResult struct {
	found       bool
	inBasin     bool
	basinNumber int
}

func outOfBounds(i, j int) bool {
	return (i < 0 || i >= length || j < 0 || j >= width)
}

func addToBasin(i, j, num int) (bool, int) {
	result := inBasinGrid[i][j]
	if result.found {
		return result.inBasin, result.basinNumber
	}
	var inBasin bool
	var basinNumberRes int

	for _, diff := range diffs {
		i_diff, j_diff := i+diff[0], j+diff[1]
		if outOfBounds(i_diff, j_diff) {
			continue
		}

		if grid[i_diff][j_diff] < num {
			for basinNumber, lowPoint := range lowPoints {
				if lowPoint[0] == i_diff && lowPoint[1] == j_diff {
					inBasinGrid[i][j] = inBasinResult{found: true, inBasin: true, basinNumber: basinNumber}
					return true, basinNumber
				} else {
					// recurse the function
					inBasin, basinNumberRes = addToBasin(i_diff, j_diff, grid[i_diff][j_diff])
					inBasinGrid[i][j] = inBasinResult{found: true, inBasin: inBasin, basinNumber: basinNumberRes}
				}
			}
		}
	}
	// inBasin will remain false if it's never been set to true
	// I'm still not 100% sure how this works
	// I wrote it, and I kinda get the logic, but still...
	return inBasin, basinNumberRes
}

func main() {

	grid = utils.ReadFileGridInts("09/input.txt")

	length = len(grid)
	width = len(grid[0])

	inBasinGrid = make([][]inBasinResult, length)
	for i := 0; i < length; i++ {
		inBasinGrid[i] = make([]inBasinResult, width)
	}

	var risk int
	diffs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i, line := range grid {
		for j, num := range line {
			low := true
			for _, diff := range diffs {
				i_diff, j_diff := i+diff[0], j+diff[1]
				if outOfBounds(i_diff, j_diff) {
					continue
				}
				if grid[i_diff][j_diff] <= num {
					// not a low point
					low = false
				}
			}
			if low {
				risk += num + 1
				lowPoints = append(lowPoints, [2]int{i, j})
			}
		}
	}

	fmt.Println(risk)

	// Part 2

	basinSizes = make([]int, len(lowPoints))

	// Adding one to every basin, as the low point is in the basin too
	for i := range basinSizes {
		basinSizes[i]++
	}

	for i, line := range grid {
		for j, num := range line {
			if num == 9 {
				continue
			}
			toAdd, pointNumber := addToBasin(i, j, num)
			if toAdd {
				basinSizes[pointNumber]++
			}
		}
	}

	sort.Ints(basinSizes)
	lastIndex := len(basinSizes)
	fmt.Println(basinSizes[lastIndex-1] * basinSizes[lastIndex-2] * basinSizes[lastIndex-3])
}
