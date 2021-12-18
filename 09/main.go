package main

import (
	"fmt"
	"strconv"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func main() {

	scanner := utils.ReadFileLines("09/input.txt")

	var grid [][]int
	// for _, line := range grid {
	// 	line = make([]int, 0)
	// }

	for scanner.Scan() {
		var list []int
		line := scanner.Text()
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			list = append(list, num)
		}
		grid = append(grid, list)
	}

	var risk int

	for i, line := range grid {
		for j, num := range line {
			low := true
			for _, diff := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				i_diff, j_diff := i+diff[0], j+diff[1]
				if i_diff < 0 || i_diff >= len(grid) || j_diff < 0 || j_diff >= len(line) {
					// out of bounds of array
					continue
				}
				if grid[i_diff][j_diff] <= num {
					// not a low point
					low = false
				}
			}
			if low {
				risk += num + 1
			}
		}
	}

	fmt.Println(risk)

}
