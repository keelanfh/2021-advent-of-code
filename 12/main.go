package main

import (
	"fmt"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

var pathMap map[string][]string
var totalPaths int

func countPaths(start string, alreadyVisited, alreadyVisitedTwice []string) {

	// if we reach the end, we know this is a valid path
	if start == "end" {
		totalPaths++
	}

	// If we've already visited this twice before, return
	if strings.ToLower(start) == start {
		for _, cave := range alreadyVisitedTwice {
			if cave == start {
				return
			}
		}
	}

	// If we've already visited this once before, add to alreadyVisitedTwice
	if strings.ToLower(start) == start {
		for _, cave := range alreadyVisited {
			if cave == start {
				// We're only allowed to visit one cave twice - start will also be in the list of visited twice
				if len(alreadyVisitedTwice) > 1 {
					return
				}
				alreadyVisitedTwice = append(alreadyVisitedTwice, start)
			}
		}
	}
	// Trying out all the possible next steps
	for _, nextStep := range pathMap[start] {
		countPaths(nextStep, append(alreadyVisited, start), alreadyVisitedTwice)
	}
}

func main() {

	scanner := utils.ReadFileLines("12/input.txt")

	pathMap = make(map[string][]string)

	for scanner.Scan() {
		pathSlice := strings.Split(scanner.Text(), "-")

		// This loop just means that we add each of the paths in both directions
		// We skip paths starting with 'end' - this stops us counting paths like A-end-A-end
		for i := 0; i <= 1; i++ {
			if pathSlice[i] != "end" {
				pathMap[pathSlice[i]] = append(pathMap[pathSlice[i]], pathSlice[1-i])
			}
		}
	}

	countPaths("start", []string{"start"}, []string{})
	fmt.Println(totalPaths)
}
