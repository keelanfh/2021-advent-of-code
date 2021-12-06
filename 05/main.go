package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var results_grid [1000][1000]int
	var part_two_grid [1000][1000]int

	for scanner.Scan() {
		// regex to match 0,9 -> 5,9

		re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
		matches := re.FindStringSubmatch(scanner.Text())

		from_x, _ := strconv.Atoi(matches[1])
		from_y, _ := strconv.Atoi(matches[2])
		to_x, _ := strconv.Atoi(matches[3])
		to_y, _ := strconv.Atoi(matches[4])

		// if (from_x == to_x) || (from_y == to_y) {
		// 	if from_x > to_x {
		// 		from_x, to_x = to_x, from_x
		// 	} else if from_y > to_y {
		// 		from_y, to_y = to_y, from_y
		// 	}
		// }

		if (from_x > to_x) || (from_y > to_y) {
			from_x, to_x = to_x, from_x
			from_y, to_y = to_y, from_y
		}

		fmt.Println(from_x, from_y, to_x, to_y)

		if from_x == to_x {
			for y := from_y; y <= to_y; y++ {
				results_grid[from_x][y]++
				part_two_grid[from_x][y]++
			}
		} else if from_y == to_y {
			for x := from_x; x <= to_x; x++ {
				results_grid[x][from_y]++
				part_two_grid[x][from_y]++
			}
		} else if (to_x - from_x) == (to_y - from_y) {
			// fmt.Println(from_x, from_y, to_x, to_y)
			for i := 0; i <= (from_x - to_x); i++ {
				part_two_grid[from_x+i][from_y+i]++
			}
		} else if (to_x - from_x) == (from_y - to_y) {
			// fmt.Println(from_x, from_y, to_x, to_y)
			for i := 0; i <= (to_x - from_x); i++ {
				part_two_grid[from_x+i][from_y-i]++
			}
		}
	}
	result := 0

	for _, row := range results_grid {
		for _, value := range row {
			if value >= 2 {
				result++
			}
		}
	}

	fmt.Println(result)

	result = 0

	for _, row := range part_two_grid {
		for _, value := range row {
			if value >= 2 {
				result++
			}
		}
	}

	fmt.Println(result)
}
