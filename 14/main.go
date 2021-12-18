package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func main() {
	scanner := utils.ReadFileLines("14/input.txt")

	var polymerTemplate []rune
	first := true
	re := regexp.MustCompile(`^([A-Z])([A-Z]) -> ([A-Z])$`)
	pairToNewPairs := make(map[[2]rune][2][2]rune)

	for scanner.Scan() {
		line := scanner.Text()
		if first {
			polymerTemplate = []rune(line)
			first = false
		} else if line != "" {
			matches := re.FindStringSubmatch(line)

			var matchesRunes [4]rune
			for i := 0; i < 4; i++ {
				matchesRunes[i] = []rune(matches[i])[0]
			}

			from := [2]rune{matchesRunes[1], matchesRunes[2]}

			to := [2][2]rune{{matchesRunes[1], matchesRunes[3]}, {matchesRunes[3], matchesRunes[2]}}

			pairToNewPairs[from] = to
		}

	}

	pairCounter := make(map[[2]rune]int)

	// reading in the starting data
	var from [2]rune
	for i := 0; i < len(polymerTemplate)-1; i++ {
		copy(from[:], polymerTemplate[i:i+2])
		pairCounter[from]++
	}

	// stepping through the required number of iterations
	for step := 1; step <= 40; step++ {
		newPairCounter := make(map[[2]rune]int)
		for current, number := range pairCounter {
			for _, newPair := range pairToNewPairs[current] {
				newPairCounter[newPair] += number
			}

		}
		pairCounter = newPairCounter
	}

	// Now we just need to go from a count of pairs to a count of letters
	letterCounter := make(map[rune]int)

	for k, v := range pairCounter {
		for _, char := range k {
			letterCounter[char] += v
		}
	}

	// Had to increase this to math.MaxInt64 because the smallest number in the array
	// is larger than math.MaxInt32
	minValue := math.MaxInt64
	maxValue := 0

	for _, value := range letterCounter {
		if value > maxValue {
			maxValue = value
		}
		if value < minValue {
			minValue = value
		}
	}

	// we're basically just halving here - but rounding up the odd numbers
	// this is because the starting and ending letters are only counted once
	// All the others are counted twice
	// The floor division will mean that adding 1 to an even number makes no difference
	fmt.Println(((maxValue+1)/2 - (minValue+1)/2))
}
