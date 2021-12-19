package main

import (
	"fmt"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

var pathMap map[string][]string
var totalPaths int
var allPaths map[string]bool

func countPaths(start string, alreadyVisited, alreadyVisitedTwice []string) int {
	fmt.Println(append(alreadyVisited, start), alreadyVisitedTwice)

	if start == "end" {
		totalPaths++
		fmt.Println("\t\t", append(alreadyVisited, start))
		allPaths[fmt.Sprint(append(alreadyVisited, start))] = true
	}

	// If we've already visited this before, return 0
	if strings.ToLower(start) == start {
		for _, cave := range alreadyVisitedTwice {
			if cave == start {
				return 0
			}
		}
	}

	// If we've already visited this before, return 0
	if strings.ToLower(start) == start {
		for _, cave := range alreadyVisited {
			if cave == start {
				alreadyVisitedTwice = append(alreadyVisitedTwice, start)
			}
		}
	}

	var totalPaths int
	for _, nextStep := range pathMap[start] {
		totalPaths += countPaths(nextStep, append(alreadyVisited, start), alreadyVisitedTwice)
		// if nextStep == "end" {
		// 	return 1
		// }
	}
	return totalPaths
}

func main() {
	allPaths = make(map[string]bool)

	scanner := utils.ReadFileLines("12/test.txt")

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
	fmt.Println(pathMap)
	fmt.Println(totalPaths)
	fmt.Println(len(allPaths))
}
