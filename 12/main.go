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
		if pathSlice[0] != "end" {
			pathMap[pathSlice[0]] = append(pathMap[pathSlice[0]], pathSlice[1])
		}
		if pathSlice[1] != "end" {
			pathMap[pathSlice[1]] = append(pathMap[pathSlice[1]], pathSlice[0])
		}
		// Need to add something here so that we don't include end- paths when declared like that
	}

	countPaths("start", []string{"start"}, []string{})
	fmt.Println(totalPaths)
}
