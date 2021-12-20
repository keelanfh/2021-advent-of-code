package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func main() {
	scanner := utils.ReadFileLines("13/input.txt")

	dots := make(map[[2]int]bool)

	re := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)
	var directions []string
	var positions []int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		} else if len(line) > 10 {
			matches := re.FindStringSubmatch(line)
			directions = append(directions, matches[1])
			position, _ := strconv.Atoi(matches[2])
			positions = append(positions, position)
		} else {
			xy := strings.Split(line, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			dots[[2]int{x, y}] = true
		}
	}

	for i, direction := range directions {
		newDots := make(map[[2]int]bool)
		position := positions[i]

		for dot, _ := range dots {
			x := dot[0]
			y := dot[1]

			if direction == "x" && x > position {
				x = position - (x - position)
			} else if direction == "y" && y > position {
				y = position - (y - position)
			}

			newDots[[2]int{x, y}] = true
		}
		dots = newDots

	}

	maxX, maxY := 0, 0
	for dot, _ := range dots {
		if dot[0] > maxX {
			maxX = dot[0]
		}
		if dot[1] > maxY {
			maxY = dot[1]
		}
	}

	var outputGrid [40][6]string

	for i := 0; i < 40; i++ {
		for j := 0; j < 6; j++ {
			outputGrid[i][j] = "."
		}
	}

	for dot, _ := range dots {
		outputGrid[dot[0]][dot[1]] = "#"
	}

	for _, line := range outputGrid {
		fmt.Println(line)
	}
}
