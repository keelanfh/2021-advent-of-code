package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func main() {
	scanner := utils.ReadFileLines("07/input.txt")

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	sliceOfStrings := strings.Split(line, ",")
	var crabs []int
	for _, s := range sliceOfStrings {
		i, _ := strconv.Atoi(s)
		crabs = append(crabs, i)
	}

	// find the max
	max := 0
	for _, i := range crabs {
		if i > max {
			max = i
		}
	}

	minFuel := math.MaxInt32
	minPartTwoFuel := math.MaxInt32

	for position := 0; position <= max; position++ {
		fuel := 0
		partTwoFuel := 0
		for _, crabPosition := range crabs {
			difference := crabPosition - position
			if difference < 0 {
				difference = -difference
			}
			fuel += difference
			// Calculating the difference - adding one to this, can't remember why
			partTwoFuel += difference * (difference + 1) / 2
		}
		if fuel < minFuel {
			minFuel = fuel
		}
		if partTwoFuel < minPartTwoFuel {
			minPartTwoFuel = partTwoFuel
		}
	}

	fmt.Println(minFuel)
	fmt.Println(minPartTwoFuel)
}
