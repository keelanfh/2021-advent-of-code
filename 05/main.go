package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func countOverlapsInGrid(grid [1000][1000]int) int {
	result := 0
	for _, row := range grid {
		for _, value := range row {
			if value >= 2 {
				result++
			}
		}
	}
	return result
}

func main() {

	scanner := utils.ReadFileLines("05/input.txt")

	var part_one_grid [1000][1000]int
	var part_two_grid [1000][1000]int

	for scanner.Scan() {
		// regex to match 0,9 -> 5,9

		re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
		matches := re.FindStringSubmatch(scanner.Text())

		from_x, _ := strconv.Atoi(matches[1])
		from_y, _ := strconv.Atoi(matches[2])
		to_x, _ := strconv.Atoi(matches[3])
		to_y, _ := strconv.Atoi(matches[4])

		// If y is the decreasing, swap them
		if from_y > to_y {
			from_x, to_x = to_x, from_x
			from_y, to_y = to_y, from_y
		}

		// If x is decreasing, swap them
		// Doing this in two steps ensures that if only one of x,y is increasing
		// i.e. for negative diagonals
		// It will be x
		if from_x > to_x {
			from_x, to_x = to_x, from_x
			from_y, to_y = to_y, from_y
		}

		// verticals
		if from_x == to_x {
			for y := from_y; y <= to_y; y++ {
				part_one_grid[from_x][y]++
				part_two_grid[from_x][y]++
			}
			// horizontals
		} else if from_y == to_y {
			for x := from_x; x <= to_x; x++ {
				part_one_grid[x][from_y]++
				part_two_grid[x][from_y]++
			}
			// positive diagonals
		} else if (to_x - from_x) == (to_y - from_y) {
			for i := 0; i <= (to_x - from_x); i++ {
				part_two_grid[from_x+i][from_y+i]++
			}
			// negative diagonals
		} else if (to_x - from_x) == (from_y - to_y) {
			for i := 0; i <= (to_x - from_x); i++ {
				part_two_grid[from_x+i][from_y-i]++
			}
		}
	}

	fmt.Println(countOverlapsInGrid(part_one_grid))

	fmt.Println(countOverlapsInGrid(part_two_grid))
}
