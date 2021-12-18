package main

import (
	"fmt"
	"strings"

	"github.com/keelanfh/2021-advent-of-code/utils"
)

func main() {
	scanner := utils.ReadFileLines("08/input.txt")

	var total int

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " | ")
		// signalPatternsString := splitLine[0]
		// signalPatternsSlice := strings.Split(signalPatternsString, " ")

		outputDigitsString := splitLine[1]
		outputDigitsSlice := strings.Split(outputDigitsString, " ")

		for _, outputDigit := range outputDigitsSlice {
			switch len(outputDigit) {
			case 2, 3, 4, 7:
				total++
			}
		}
	}
	fmt.Println(total)
}
